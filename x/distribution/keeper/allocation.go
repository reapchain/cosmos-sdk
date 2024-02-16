package keeper

import (
	"fmt"
	"math/big"

	abci "github.com/reapchain/reapchain-core/abci/types"

	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/x/distribution/types"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
)

const (
	// reward rate
	StandingMemberRewardRate = "0.1"
	SteeringMemberRewardRate = "0.2"
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

	// fetch and clear the collected fees for distribution, since this is
	// called in BeginBlock, collected fees will be from the previous block
	// (and distributed to the previous proposer)
	feeCollector := k.authKeeper.GetModuleAccount(ctx, k.feeCollectorName)
	feesCollectedInt := k.bankKeeper.GetAllBalances(ctx, feeCollector.GetAddress())

	if feesCollectedInt.IsZero() {
		return
	}

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
	/*
		No rewards for proposer(2022.08.04)


		// calculate fraction votes
		previousFractionVotes := sdk.NewDec(sumPreviousPrecommitPower).Quo(sdk.NewDec(totalPreviousPower))

		// calculate previous proposer reward
		baseProposerReward := k.GetBaseProposerReward(ctx)
		bonusProposerReward := k.GetBonusProposerReward(ctx)
		proposerMultiplier := baseProposerReward.Add(bonusProposerReward.MulTruncate(previousFractionVotes))
		proposerReward := feesCollected.MulDecTruncate(proposerMultiplier)
	*/
	remaining := feesCollected

	// make reward targets

	// 1. steeringMemberCandidatesLived is a list of steering member candidate that was operating in the previous round.
	steeringMemberCandidatesLived := make([]stakingtypes.ValidatorI, len(vrfList.VrfCheckList))
	i := 0
	for _, smc := range vrfList.VrfCheckList {
		if smc.IsVrfTransmission {
			val := k.stakingKeeper.ValidatorByConsAddr(ctx, smc.GetSteeringMemberCandidateAddress())
			if val == nil {
				k.Logger(ctx).Info(fmt.Sprintf("Not exist Steering Member validator: %s", smc.GetSteeringMemberCandidateAddress()))
				continue
			}
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
		if validator.GetType() == stakingtypes.ValidatorTypeStanding {
			standingMembers = append(standingMembers, validator)
		} else if validator.GetType() == stakingtypes.ValidatorTypeSteering {
			steeringMembers = append(steeringMembers, validator)
		}
	}

	// 4. allValidators is a list of validator in the Block of corresponding height.
	allValidators := append(standingMembers, steeringMemberCandidatesLived...)

	// pay previous proposer
	proposerValidator := k.stakingKeeper.ValidatorByConsAddr(ctx, previousProposer)

	if proposerValidator != nil {
		/* No rewards for proposer(2022.08.04)

		ctx.EventManager().EmitEvent(
			sdk.NewEvent(
				types.EventTypeProposerReward,
				sdk.NewAttribute(sdk.AttributeKeyAmount, proposerReward.String()),
				sdk.NewAttribute(types.AttributeKeyValidator, proposerValidator.GetOperator().String()),
			),
		)

		k.AllocateTokensToValidator(ctx, proposerValidator, proposerReward)
		remaining = remaining.Sub(proposerReward)
		*/
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
	// TODO: need to set genesis params
	standingMemberRewardPercent := k.GetStandingMemberRewardRate(ctx)
	steeringMemberRewardPercent := k.GetSteeringMemberRewardRate(ctx)

	fmt.Println("standingMemberRewardPercent: ", standingMemberRewardPercent)
	fmt.Println("steeringMemberRewardPercent: ", steeringMemberRewardPercent)

	// calculate standing member rewards
	standingMemberReward := feesCollected.MulDecTruncate(standingMemberRewardPercent)
	// calculate steering member rewards
	steeringMemberReward := feesCollected.MulDecTruncate(steeringMemberRewardPercent)

	// we'll set zero.
	communityTax := k.GetCommunityTax(ctx)

	// calculate all validator(alive) rewards
	// No rewards for proposer(2022.08.04)
	//voteMultiplier := sdk.OneDec().Sub(proposerMultiplier).Sub(standingMemberRewardPercent).Sub(steeringMemberRewardPercent).Sub(communityTax)
	voteMultiplier := sdk.OneDec().Sub(standingMemberRewardPercent).Sub(steeringMemberRewardPercent).Sub(communityTax)
	allValidatorReward := feesCollected.MulDecTruncate(voteMultiplier)

	// allocate tokens to validators who lived node
	totalValidator := sdk.NewDec(int64(len(allValidators)))
	totalStandingMember := sdk.NewDec(int64(len(standingMembers)))
	totalStandingMemberRates := totalStandingMember.QuoTruncate(totalValidator)

	standingMemberReward2 := allValidatorReward.MulDecTruncate(totalStandingMemberRates)
	steeringMemberReward2 := allValidatorReward.MulDecTruncate(sdk.OneDec().Sub(totalStandingMemberRates))

	// allocate tokens to Standing Members
	// calculate total standing members power
	totalPowerStandingMember := int64(0)
	for _, val := range standingMembers {
		totalPowerStandingMember += val.GetConsensusPower(sdk.DefaultPowerReduction)
	}

	if totalPowerStandingMember > 0 {

		for _, val := range standingMembers {
			// calculate each validator's power and rewards rate.
			valPower := sdk.NewDecFromBigInt(big.NewInt(val.GetConsensusPower(sdk.DefaultPowerReduction)))
			rewardRate := valPower.QuoTruncate(sdk.NewDec(totalPowerStandingMember))

			// calculate each reward from standing member's total rewards
			rewards := standingMemberReward.MulDecTruncate(rewardRate)
			rewards2 := standingMemberReward2.MulDecTruncate(rewardRate)

			totalRewards := rewards.Add(rewards2...)

			k.AllocateTokensToValidator(ctx, val, totalRewards)

			remaining = remaining.Sub(totalRewards)
		}
	}

	// allocate tokens to Steering Members
	totalSteeringMember := sdk.NewInt(int64(len(steeringMembers)))
	for _, val := range steeringMembers {
		rewards := steeringMemberReward.MulDecTruncate(sdk.OneDec().QuoTruncate(sdk.NewDecFromInt(totalSteeringMember)))

		k.AllocateTokensToValidator(ctx, val, rewards)

		remaining = remaining.Sub(rewards)
	}

	// allocate tokens to Steering Members candidates lived
	totalSteeringMemberCandidatesLived := sdk.NewInt(int64(len(steeringMemberCandidatesLived)))
	for _, val := range steeringMemberCandidatesLived {
		rewards := steeringMemberReward2.MulDecTruncate(sdk.OneDec().QuoTruncate(sdk.NewDecFromInt(totalSteeringMemberCandidatesLived)))

		k.AllocateTokensToValidator(ctx, val, rewards)

		remaining = remaining.Sub(rewards)
	}

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
	fmt.Println("valAddress: ", val.GetOperator())
	fmt.Println("tokens: ", tokens)
	fmt.Println("commission: ", commission)
	fmt.Println("shared: ", shared)
	fmt.Println("currentRewards.Rewards: ", currentRewards.Rewards)
	fmt.Println("outstanding.Rewards: ", outstanding.Rewards)
	fmt.Println("####################################################################")

}
