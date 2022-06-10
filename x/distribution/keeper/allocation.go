package keeper

import (
	"fmt"

	abci "github.com/reapchain/reapchain-core/abci/types"

	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
)

// AllocateTokens handles distribution of the collected fees
// bondedVotes is a list of (validator address, validator voted on last block flag) for all
// validators in the bonded set.
func (k Keeper) AllocateTokens(
	ctx sdk.Context, sumPreviousPrecommitPower, totalPreviousPower int64,
	previousProposer sdk.ConsAddress, bondedVotes []abci.VoteInfo,
	vrfList abci.VrfCheckList,
) {

	logger := k.Logger(ctx)

	// make reward targets

	// 1. steeringMemberCandidatesLived is a list of steering member candidate that was operating in the previous round.
	steeringMemberCandidatesLived := make([]stakingtypes.ValidatorI, len(vrfList.VrfCheckList))
	i := 0
	for _, smc := range vrfList.VrfCheckList {
		if smc.IsVrfTransmission {
			val := k.stakingKeeper.ValidatorByConsAddr(ctx, smc.GetSteeringMemberCandidateAddress())
			steeringMemberCandidatesLived[i] = val
			i++
		}
	}
	steeringMemberCandidatesLived = steeringMemberCandidatesLived[:i]

	// 2. standingMembers is a list of validator(standing member) in the Block of corresponding height.
	var standingMembers []stakingtypes.ValidatorI

	// 3. steeringMembers is a list of validator(steering member) in the Block of corresponding height.
	var steeringMembers []stakingtypes.ValidatorI

	for _, voteInfo := range bondedVotes {
		validator := k.stakingKeeper.ValidatorByConsAddr(ctx, voteInfo.Validator.GetAddress())
		if validator.GetType() == "standing" {
			standingMembers = append(standingMembers, validator)
		} else if validator.GetType() == "steering" {
			steeringMembers = append(steeringMembers, validator)
		}
	}

	// 4. allValidators is a list of validator in the Block of corresponding height.
	allValidators := append(standingMembers, steeringMemberCandidatesLived...)

	//////////////////////////////////////////////////////////////////////////////////////////////////
	fmt.Println("######## validator info ########")
	fmt.Println("steeringMemberCandidatesLived: ", len(steeringMemberCandidatesLived))
	fmt.Println("standingMembers: ", len(standingMembers))
	fmt.Println("steeringMembers: ", len(steeringMembers))
	fmt.Println("allValidators: ", len(allValidators))
	//////////////////////////////////////////////////////////////////////////////////////////////////

	// fetch and clear the collected fees for distribution, since this is
	// called in BeginBlock, collected fees will be from the previous block
	// (and distributed to the previous proposer)
	feeCollector := k.authKeeper.GetModuleAccount(ctx, k.feeCollectorName)
	feesCollectedInt := k.bankKeeper.GetAllBalances(ctx, feeCollector.GetAddress())

	// For distribution test
	//if feesCollectedInt.IsAllGT(sdk.NewCoins(sdk.NewInt64Coin("areap", 1000000))) {
	//	feesCollectedInt = sdk.NewCoins(sdk.NewInt64Coin("areap", 1000000))
	//}

	feesCollected := sdk.NewDecCoinsFromCoins(feesCollectedInt...)

	// transfer collected fees to the distribution module account
	err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, k.feeCollectorName, types.ModuleName, feesCollectedInt)
	if err != nil {
		panic(err)
	}

	// temporary workaround to keep CanWithdrawInvariant happy
	// general discussions here: https://github.com/reapchain/cosmos-sdk/issues/2906#issuecomment-441867634
	feePool := k.GetFeePool(ctx)
	if totalPreviousPower == 0 {
		feePool.CommunityPool = feePool.CommunityPool.Add(feesCollected...)
		k.SetFeePool(ctx, feePool)
		return
	}

	// calculate fraction votes
	previousFractionVotes := sdk.NewDec(sumPreviousPrecommitPower).Quo(sdk.NewDec(totalPreviousPower))

	// calculate previous proposer reward
	baseProposerReward := k.GetBaseProposerReward(ctx)
	bonusProposerReward := k.GetBonusProposerReward(ctx)
	proposerMultiplier := baseProposerReward.Add(bonusProposerReward.MulTruncate(previousFractionVotes))
	proposerReward := feesCollected.MulDecTruncate(proposerMultiplier)

	// pay previous proposer
	remaining := feesCollected
	proposerValidator := k.stakingKeeper.ValidatorByConsAddr(ctx, previousProposer)

	if proposerValidator != nil {
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeProposerReward,
				sdk.NewAttribute(sdk.AttributeKeyAmount, proposerReward.String()),
				sdk.NewAttribute(types.AttributeKeyValidator, proposerValidator.GetOperator().String()),
			),
		)

		k.AllocateTokensToValidator(ctx, proposerValidator, proposerReward)
		remaining = remaining.Sub(proposerReward)
	} else {
		// previous proposer can be unknown if say, the unbonding period is 1 block, so
		// e.g. a validator undelegates at block X, it's removed entirely by
		// block X+1's endblock, then X+2 we need to refer to the previous
		// proposer for X+1, but we've forgotten about them.
		logger.Error(fmt.Sprintf(
			"WARNING: Attempt to allocate proposer rewards to unknown proposer %s. "+
				"This should happen only if the proposer unbonded completely within a single block, "+
				"which generally should not happen except in exceptional circumstances (or fuzz testing). "+
				"We recommend you investigate immediately.",
			previousProposer.String()))
	}

	// reward rate.
	standingMemberRewardPercent, _ := sdk.NewDecFromStr("0.1") // 10%
	steeringMemberRewardPercent, _ := sdk.NewDecFromStr("0.2") // 20%

	// calculate standing member rewards
	standingMemberReward := feesCollected.MulDecTruncate(standingMemberRewardPercent)

	// calculate steering member rewards
	steeringMemberReward := feesCollected.MulDecTruncate(steeringMemberRewardPercent)

	// we'll set zero.
	communityTax := k.GetCommunityTax(ctx)

	// calculate all validator(alive) rewards
	voteMultiplier := sdk.OneDec().Sub(proposerMultiplier).Sub(standingMemberRewardPercent).Sub(steeringMemberRewardPercent).Sub(communityTax)
	allValidatorReward := feesCollected.MulDecTruncate(voteMultiplier)

	//remaining = remaining.Sub(standingMemberReward).Sub(steeringMemberReward).Sub(allValidatorReward)

	//////////////////////////////////////////////////////////////////////////////////////////////////
	fmt.Println("############ Rewards ############")
	fmt.Println("feesCollected: ", feesCollected)

	fmt.Println("baseProposerReward: ", baseProposerReward)
	fmt.Println("bonusProposerReward: ", bonusProposerReward)
	fmt.Println("proposerMultiplier: ", proposerMultiplier)
	fmt.Println("proposerReward: ", proposerReward)

	fmt.Println("communityTax: ", communityTax)
	fmt.Println("voteMultiplier: ", voteMultiplier)

	fmt.Println("standingMemberReward: ", standingMemberReward)
	fmt.Println("steeringMemberReward: ", steeringMemberReward)
	fmt.Println("allValidatorReward: ", allValidatorReward)
	//fmt.Println("remain: ", remaining)
	//////////////////////////////////////////////////////////////////////////////////////////////////

	// allocate tokens to Standing Members
	totalStandingMember := sdk.NewInt(int64(len(standingMembers)))
	for _, val := range standingMembers {
		reward := standingMemberReward.MulDecTruncate(sdk.OneDec().QuoTruncate(sdk.NewDecFromInt(totalStandingMember)))

		// For debug
		addr, _ := val.GetConsAddr()
		fmt.Println("Standing 1: ", addr, reward)

		k.AllocateTokensToValidator(ctx, val, reward)
		remaining = remaining.Sub(reward)
	}

	fmt.Println("remain: ", remaining)

	// allocate tokens to Steering Members
	totalSteeringMember := sdk.NewInt(int64(len(steeringMembers)))
	for _, val := range steeringMembers {
		reward := steeringMemberReward.MulDecTruncate(sdk.OneDec().QuoTruncate(sdk.NewDecFromInt(totalSteeringMember)))

		// For debug
		addr, _ := val.GetConsAddr()
		fmt.Println("Steering 1: ", addr, reward)

		k.AllocateTokensToValidator(ctx, val, reward)
		remaining = remaining.Sub(reward)
	}

	fmt.Println("remain: ", remaining)

	// allocate tokens to validators who lived node
	totalValidator := sdk.NewInt(int64(len(allValidators)))
	for _, val := range allValidators {
		reward := allValidatorReward.MulDecTruncate(sdk.OneDec().QuoTruncate(sdk.NewDecFromInt(totalValidator)))

		// For debug
		addr, _ := val.GetConsAddr()
		fmt.Println("Validator 1: ", addr, reward)

		k.AllocateTokensToValidator(ctx, val, reward)
		remaining = remaining.Sub(reward)
	}

	fmt.Println("remain: ", remaining)

	/*
		// allocate tokens proportionally to voting power
		// TODO consider parallelizing later, ref https://github.com/reapchain/cosmos-sdk/pull/3099#discussion_r246276376
		for _, vote := range bondedVotes {
			validator := k.stakingKeeper.ValidatorByConsAddr(ctx, vote.Validator.Address)

			// TODO consider microslashing for missing votes.
			// ref https://github.com/reapchain/cosmos-sdk/issues/2525#issuecomment-430838701
			powerFraction := sdk.NewDec(vote.Validator.Power).QuoTruncate(sdk.NewDec(totalPreviousPower))
			reward := feesCollected.MulDecTruncate(voteMultiplier).MulDecTruncate(powerFraction)
			k.AllocateTokensToValidator(ctx, validator, reward)
			remaining = remaining.Sub(reward)
		}
	*/
	// allocate community funding
	feePool.CommunityPool = feePool.CommunityPool.Add(remaining...)
	k.SetFeePool(ctx, feePool)
}

// AllocateTokensToValidator allocate tokens to a particular validator, splitting according to commission
func (k Keeper) AllocateTokensToValidator(ctx sdk.Context, val stakingtypes.ValidatorI, tokens sdk.DecCoins) {
	// split tokens between validator and delegators according to commission
	commission := tokens.MulDec(val.GetCommission())
	shared := tokens.Sub(commission)

	// update current commission
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeCommission,
			sdk.NewAttribute(sdk.AttributeKeyAmount, commission.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, val.GetOperator().String()),
		),
	)
	currentCommission := k.GetValidatorAccumulatedCommission(ctx, val.GetOperator())
	currentCommission.Commission = currentCommission.Commission.Add(commission...)
	k.SetValidatorAccumulatedCommission(ctx, val.GetOperator(), currentCommission)

	// update current rewards
	currentRewards := k.GetValidatorCurrentRewards(ctx, val.GetOperator())
	currentRewards.Rewards = currentRewards.Rewards.Add(shared...)
	k.SetValidatorCurrentRewards(ctx, val.GetOperator(), currentRewards)

	// update outstanding rewards
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeRewards,
			sdk.NewAttribute(sdk.AttributeKeyAmount, tokens.String()),
			sdk.NewAttribute(types.AttributeKeyValidator, val.GetOperator().String()),
		),
	)
	outstanding := k.GetValidatorOutstandingRewards(ctx, val.GetOperator())
	outstanding.Rewards = outstanding.Rewards.Add(tokens...)
	k.SetValidatorOutstandingRewards(ctx, val.GetOperator(), outstanding)

	fmt.Println("######################## Allocate Validator ########################")
	fmt.Println("tokens: ", tokens)
	fmt.Println("commission: ", commission)
	fmt.Println("shared: ", shared)
	fmt.Println("currentRewards.Rewards: ", currentRewards.Rewards)
	fmt.Println("outstanding.Rewards: ", outstanding.Rewards)
	fmt.Println("######################## Allocate Validator ########################")

}
