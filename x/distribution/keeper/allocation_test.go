package keeper_test

import (
	"github.com/reapchain/cosmos-sdk/x/staking"
	"testing"

	abci "github.com/reapchain/reapchain-core/abci/types"
	tmproto "github.com/reapchain/reapchain-core/proto/reapchain-core/types"
	"github.com/stretchr/testify/require"

	"github.com/reapchain/cosmos-sdk/simapp"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/x/auth/types"
	"github.com/reapchain/cosmos-sdk/x/staking/teststaking"
	stakingtypes "github.com/reapchain/cosmos-sdk/x/staking/types"
)

const cntValidators, cntStandingMembers, cntSteeringMembers = 29, 14, 15

var (
	app      *simapp.SimApp
	ctx      sdk.Context
	addrs    []sdk.AccAddress
	valAddrs []sdk.ValAddress
)

func init() {
}

/*
func TestAllocateTokensToValidatorWithCommission(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrs(app, ctx, 3, sdk.NewInt(45000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addrs)
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	// create validator with 50% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	tstaking.CreateValidator(sdk.ValAddress(addrs[0]), valConsPk1, sdk.NewInt(44000000), true, stakingtypes.ValidatorTypeStanding)
	val := app.StakingKeeper.Validator(ctx, valAddrs[0])

	// allocate tokens
	tokens := sdk.DecCoins{
		{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(10)},
	}
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)

	// check commission
	expected := sdk.DecCoins{
		{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(5)},
	}
	require.Equal(t, expected, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, val.GetOperator()).Commission)

	// check current rewards
	require.Equal(t, expected, app.DistrKeeper.GetValidatorCurrentRewards(ctx, val.GetOperator()).Rewards)
}

func TestAllocateTokensToManyValidators(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrs(app, ctx, 5, sdk.NewInt(45000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addrs)
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	// create with 50% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[0], valConsPk1, sdk.NewInt(44000000), true, stakingtypes.ValidatorTypeStanding)

	// create validator with 0% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[1], valConsPk2, sdk.NewInt(100000), true, stakingtypes.ValidatorTypeSteering)

	// create validator with 0% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[2], valConsPk3, sdk.NewInt(100000), true, stakingtypes.ValidatorTypeSteering)

	// create validator with 0% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[3], valConsPk4, sdk.NewInt(100000), true, stakingtypes.ValidatorTypeSteering)

	// create validator with 0% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[4], valConsPk5, sdk.NewInt(100000), true, stakingtypes.ValidatorTypeSteering)

	abciValA := abci.Validator{
		Address: valConsPk1.Address(),
		Power:   10,
	}
	abciValB := abci.Validator{
		Address: valConsPk2.Address(),
		Power:   10,
	}
	abciValC := abci.Validator{
		Address: valConsPk3.Address(),
		Power:   10,
	}
	abciValD := abci.Validator{
		Address: valConsPk4.Address(),
		Power:   10,
	}
	abciValE := abci.Validator{
		Address: valConsPk5.Address(),
		Power:   10,
	}

	// assert initial state: zero outstanding rewards, zero community pool, zero commission, zero current rewards
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[1]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[2]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[3]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[4]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetFeePool(ctx).CommunityPool.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[1]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[2]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[3]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[4]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[0]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[1]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[2]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[3]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[4]).Rewards.IsZero())

	// allocate tokens as if both had voted and second was proposer
	fees := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)))
	feeCollector := app.AccountKeeper.GetModuleAccount(ctx, types.FeeCollectorName)
	require.NotNil(t, feeCollector)

	// fund fee collector
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, feeCollector.GetName(), fees))

	app.AccountKeeper.SetAccount(ctx, feeCollector)

	votes := []abci.VoteInfo{
		{
			Validator:       abciValA,
			SignedLastBlock: true,
		},
		{
			Validator:       abciValB,
			SignedLastBlock: true,
		},
		{
			Validator:       abciValC,
			SignedLastBlock: true,
		},
		{
			Validator:       abciValD,
			SignedLastBlock: true,
		},
		{
			Validator:       abciValE,
			SignedLastBlock: true,
		},
	}
	var vrfList abci.VrfCheckList
	vrfList.VrfCheckList = make([]*abci.VrfCheck, 4)
	vrfList.VrfCheckList[0] = &abci.VrfCheck{
		SteeringMemberCandidateAddress: valConsAddr3,
		IsVrfTransmission:              true,
	}
	vrfList.VrfCheckList[1] = &abci.VrfCheck{
		SteeringMemberCandidateAddress: valConsAddr4,
		IsVrfTransmission:              true,
	}
	vrfList.VrfCheckList[2] = &abci.VrfCheck{
		SteeringMemberCandidateAddress: valConsAddr5,
		IsVrfTransmission:              true,
	}
	vrfList.VrfCheckList[3] = &abci.VrfCheck{
		SteeringMemberCandidateAddress: valConsAddr2,
		IsVrfTransmission:              true,
	}

	app.DistrKeeper.AllocateTokens(ctx, 200, 200, valConsAddr1, votes, vrfList)

	// (TotalRewards * proposer_rate) + (TotalRewards * bonus_rate) + (TotalRewards * standing_rate)/standingCount +  + (TotalRewards * AllValidator_rete)/AllValidatorCount
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(280, 1)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards)
	//
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(180, 1)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[1]).Rewards)
	//
	require.Equal(t, sdk.DecCoins(sdk.DecCoins(nil)), app.DistrKeeper.GetFeePool(ctx).CommunityPool)
	// 28 * 0.5
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(140, 1)}}, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission)
	// zero commission for second proposer
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[1]).Commission.IsZero())
	//
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(140, 1)}}, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[0]).Rewards)
	// proposer reward + staking.proportional for second proposer = (5 % + 0.5 * (93%)) * 100 = 51.5
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(180, 1)}}, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[1]).Rewards)
}

func TestAllocateTokensTruncation(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrs(app, ctx, 5, sdk.NewInt(100000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addrs)
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	// create validator with 10% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(1, 1), sdk.NewDecWithPrec(1, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[0], valConsPk1, sdk.NewInt(44000000), true, stakingtypes.ValidatorTypeStanding)

	// create second validator with 100% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(10, 1), sdk.NewDecWithPrec(10, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[1], valConsPk2, sdk.NewInt(100000), true, stakingtypes.ValidatorTypeSteering)

	// create third validator with 100% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(10, 1), sdk.NewDecWithPrec(10, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[2], valConsPk3, sdk.NewInt(100000), true, stakingtypes.ValidatorTypeSteering)

	// create third validator with 100% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(10, 1), sdk.NewDecWithPrec(10, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[3], valConsPk4, sdk.NewInt(100000), true, stakingtypes.ValidatorTypeSteering)

	// create third validator with 100% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(10, 1), sdk.NewDecWithPrec(10, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[4], valConsPk5, sdk.NewInt(100000), true, stakingtypes.ValidatorTypeSteering)

	abciValA := abci.Validator{
		Address: valConsPk1.Address(),
		Power:   10,
	}
	abciValB := abci.Validator{
		Address: valConsPk2.Address(),
		Power:   10,
	}
	abciValС := abci.Validator{
		Address: valConsPk3.Address(),
		Power:   10,
	}
	abciValD := abci.Validator{
		Address: valConsPk4.Address(),
		Power:   10,
	}
	abciValE := abci.Validator{
		Address: valConsPk5.Address(),
		Power:   10,
	}

	// assert initial state: zero outstanding rewards, zero community pool, zero commission, zero current rewards
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[1]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[2]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[3]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[4]).Rewards.IsZero())

	require.True(t, app.DistrKeeper.GetFeePool(ctx).CommunityPool.IsZero())

	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[1]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[2]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[3]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[4]).Commission.IsZero())

	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[0]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[1]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[2]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[3]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[4]).Rewards.IsZero())

	fees := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(100)))

	feeCollector := app.AccountKeeper.GetModuleAccount(ctx, types.FeeCollectorName)
	require.NotNil(t, feeCollector)

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, feeCollector.GetName(), fees))

	app.AccountKeeper.SetAccount(ctx, feeCollector)

	votes := []abci.VoteInfo{
		{
			Validator:       abciValA,
			SignedLastBlock: true,
		},
		{
			Validator:       abciValB,
			SignedLastBlock: true,
		},
		{
			Validator:       abciValС,
			SignedLastBlock: true,
		},
		{
			Validator:       abciValD,
			SignedLastBlock: true,
		},
		{
			Validator:       abciValE,
			SignedLastBlock: true,
		},
	}

	var vrfList abci.VrfCheckList
	vrfList.VrfCheckList = make([]*abci.VrfCheck, 4)
	vrfList.VrfCheckList[0] = &abci.VrfCheck{
		SteeringMemberCandidateAddress: valConsAddr2,
		IsVrfTransmission:              true,
	}
	vrfList.VrfCheckList[1] = &abci.VrfCheck{
		SteeringMemberCandidateAddress: valConsAddr3,
		IsVrfTransmission:              true,
	}
	vrfList.VrfCheckList[2] = &abci.VrfCheck{
		SteeringMemberCandidateAddress: valConsAddr4,
		IsVrfTransmission:              true,
	}
	vrfList.VrfCheckList[3] = &abci.VrfCheck{
		SteeringMemberCandidateAddress: valConsAddr5,
		IsVrfTransmission:              true,
	}

	app.DistrKeeper.AllocateTokens(ctx, 30, 30, sdk.ConsAddress(valConsPk1.Address()), votes, vrfList)

	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards.IsValid())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[1]).Rewards.IsValid())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[2]).Rewards.IsValid())

	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(280, 1)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards)
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(252, 1)}}, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[0]).Rewards)
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(28, 1)}}, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission)

	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(180, 1)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[1]).Rewards)
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(180, 1)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[2]).Rewards)

}
*/

// 0. init test
func InitTest(t *testing.T) {

	app = simapp.Setup(false)
	ctx = app.BaseApp.NewContext(false, tmproto.Header{})

	addrs = simapp.AddTestAddrs(app, ctx, cntValidators, sdk.NewInt(100000000))
	valAddrs = simapp.ConvertAddrsToValAddrs(addrs)

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	for i := 0; i < cntValidators; i++ {
		if i < cntStandingMembers {
			// create standing validator with 10% commission
			tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(1, 1), sdk.NewDecWithPrec(1, 1), sdk.NewDec(0))
			tstaking.CreateValidator(valAddrs[i], PKS[i], sdk.NewInt(44000000), true, stakingtypes.ValidatorTypeStanding)
		} else {
			// create steering validator with 100% commission
			tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDec(0), sdk.NewDec(0), sdk.NewDec(0))
			tstaking.CreateValidator(valAddrs[i], PKS[i], sdk.NewInt(100000), true, stakingtypes.ValidatorTypeSteering)
		}

	}

	for i := 0; i < cntValidators; i++ {
		// assert initial state: zero outstanding rewards, zero community pool, zero commission, zero current rewards
		require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards.IsZero())
		require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[i]).Commission.IsZero())
		require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[i]).Rewards.IsZero())
	}

	require.True(t, app.DistrKeeper.GetFeePool(ctx).CommunityPool.IsZero())
}

// unbond validator
func unbondValidator(t *testing.T, no int, amt sdk.Int) {
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	// unbond validator total self-delegations (which should jail the validator)
	valAcc := sdk.AccAddress(valAddrs[no])
	tstaking.Undelegate(valAcc, valAddrs[no], amt, true)

	_, err := app.StakingKeeper.CompleteUnbonding(ctx, sdk.AccAddress(valAddrs[no]), valAddrs[no])
	require.Nil(t, err, "expected complete unbonding validator to be ok, got: %v", err)

}

// 1. basic distribution test
func TestBasicAllocateTokens(t *testing.T) {
	InitTest(t)

	fees := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000)))

	feeCollector := app.AccountKeeper.GetModuleAccount(ctx, types.FeeCollectorName)
	require.NotNil(t, feeCollector)

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, feeCollector.GetName(), fees))

	app.AccountKeeper.SetAccount(ctx, feeCollector)

	var abciVal [cntValidators]abci.Validator
	var votes []abci.VoteInfo
	votes = make([]abci.VoteInfo, cntValidators)

	for i := 0; i < cntValidators; i++ {
		abciVal[i] = abci.Validator{
			Address: PKS[i].Address(),
			Power:   10,
		}

		votes[i] = abci.VoteInfo{
			Validator:       abciVal[i],
			SignedLastBlock: true,
		}
	}
	var vrfList abci.VrfCheckList
	vrfList.VrfCheckList = make([]*abci.VrfCheck, cntSteeringMembers)

	for i := 0; i < cntSteeringMembers; i++ {
		vrfList.VrfCheckList[i] = &abci.VrfCheck{
			SteeringMemberCandidateAddress: sdk.ConsAddress(PKS[cntStandingMembers+i].Address()),
			IsVrfTransmission:              true,
		}
	}

	app.DistrKeeper.AllocateTokens(ctx, 30, 30, sdk.ConsAddress(PKS[0].Address()), votes, vrfList)

	for i := 0; i < cntValidators; i++ {
		require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards.IsValid())
	}

	//standing + proposer(1%) + bonus(4%)
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3128078817733990130, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards)
	//standing
	for i := 1; i < cntStandingMembers; i++ {
		require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3128078817733990130, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards)
	}
	//steering
	for i := cntStandingMembers; i < cntValidators; i++ {
		require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3747126436781609170, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards)
	}
}

// 2. unbond standing validator
func TestAllocateTokensAfterUnbondStanding(t *testing.T) {
	InitTest(t)
	unbondValidator(t, 2, sdk.NewInt(44000000))

	allVals := app.StakingKeeper.GetAllValidators(ctx)
	require.Equal(t, 28, len(allVals))

	fees := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000)))

	feeCollector := app.AccountKeeper.GetModuleAccount(ctx, types.FeeCollectorName)
	require.NotNil(t, feeCollector)

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, feeCollector.GetName(), fees))

	app.AccountKeeper.SetAccount(ctx, feeCollector)

	var abciVal [cntValidators]abci.Validator
	var votes []abci.VoteInfo
	votes = make([]abci.VoteInfo, cntValidators-1)

	j := 0
	for i := 0; i < cntValidators; i++ {
		abciVal[i] = abci.Validator{
			Address: PKS[i].Address(),
			Power:   10,
		}

		if i != 2 {
			votes[j] = abci.VoteInfo{
				Validator:       abciVal[i],
				SignedLastBlock: true,
			}
			j++
		}
	}
	var vrfList abci.VrfCheckList
	vrfList.VrfCheckList = make([]*abci.VrfCheck, cntSteeringMembers)

	for i := 0; i < cntSteeringMembers; i++ {
		vrfList.VrfCheckList[i] = &abci.VrfCheck{
			SteeringMemberCandidateAddress: sdk.ConsAddress(PKS[cntStandingMembers+i].Address()),
			IsVrfTransmission:              true,
		}
	}

	app.DistrKeeper.AllocateTokens(ctx, 30, 30, sdk.ConsAddress(PKS[0].Address()), votes, vrfList)

	for i := 0; i < cntValidators; i++ {
		require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards.IsValid())
	}

	//standing
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3269230769230769210, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards)
	//standing
	for i := 1; i < cntStandingMembers && i != 2; i++ {
		require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3269230769230769210, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards)
	}
	//steering
	for i := cntStandingMembers; i < cntValidators; i++ {
		require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3833333333333333300, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards)
	}

}

// 3. unbond steering validator
func TestAllocateTokensAfterUnbondSteering(t *testing.T) {
	InitTest(t)
	unbondValidator(t, 15, sdk.NewInt(100000))

	allVals := app.StakingKeeper.GetAllValidators(ctx)
	require.Equal(t, 28, len(allVals))

	fees := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000)))

	feeCollector := app.AccountKeeper.GetModuleAccount(ctx, types.FeeCollectorName)
	require.NotNil(t, feeCollector)

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, feeCollector.GetName(), fees))

	app.AccountKeeper.SetAccount(ctx, feeCollector)

	var abciVal [cntValidators]abci.Validator
	var votes []abci.VoteInfo
	votes = make([]abci.VoteInfo, cntValidators-1)

	j := 0
	for i := 0; i < cntValidators; i++ {
		abciVal[i] = abci.Validator{
			Address: PKS[i].Address(),
			Power:   10,
		}

		if i != 15 {
			votes[j] = abci.VoteInfo{
				Validator:       abciVal[i],
				SignedLastBlock: true,
			}
			j++
		}
	}
	var vrfList abci.VrfCheckList
	vrfList.VrfCheckList = make([]*abci.VrfCheck, cntSteeringMembers-1)

	j = 0
	for i := 0; i < cntSteeringMembers; i++ {
		if i != 1 {
			vrfList.VrfCheckList[j] = &abci.VrfCheck{
				SteeringMemberCandidateAddress: sdk.ConsAddress(PKS[cntStandingMembers+i].Address()),
				IsVrfTransmission:              true,
			}
			j++
		}
	}

	app.DistrKeeper.AllocateTokens(ctx, 30, 30, sdk.ConsAddress(PKS[0].Address()), votes, vrfList)

	for i := 0; i < cntValidators; i++ {
		require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards.IsValid())
	}

	//standing
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3214285714285714260, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards)
	//standing
	for i := 1; i < cntStandingMembers; i++ {
		require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3214285714285714260, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards)
	}
	//steering
	for i := cntStandingMembers; i < cntValidators && i != 15; i++ {
		require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3928571428571428540, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards)
	}

}

// 3. delegate rewards test
func TestAllocateTokensDelegate(t *testing.T) {
	InitTest(t)

	// end block to bond validator and start new block
	staking.EndBlocker(ctx, app.StakingKeeper)
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)

	// delegate to standing #1
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	valAcc := sdk.AccAddress(valAddrs[15])
	tstaking.Delegate(valAcc, valAddrs[1], sdk.NewInt(4400000))
	del := app.StakingKeeper.Delegation(ctx, valAcc, valAddrs[1])

	require.Equal(t, sdk.NewDec(4400000), del.GetShares())

	val := app.StakingKeeper.Validator(ctx, valAddrs[1])

	fees := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10000)))

	feeCollector := app.AccountKeeper.GetModuleAccount(ctx, types.FeeCollectorName)
	require.NotNil(t, feeCollector)

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, feeCollector.GetName(), fees))

	app.AccountKeeper.SetAccount(ctx, feeCollector)

	var abciVal [cntValidators]abci.Validator
	var votes []abci.VoteInfo
	votes = make([]abci.VoteInfo, cntValidators)

	for i := 0; i < cntValidators; i++ {
		abciVal[i] = abci.Validator{
			Address: PKS[i].Address(),
			Power:   10,
		}

		votes[i] = abci.VoteInfo{
			Validator:       abciVal[i],
			SignedLastBlock: true,
		}
	}
	var vrfList abci.VrfCheckList
	vrfList.VrfCheckList = make([]*abci.VrfCheck, cntSteeringMembers)

	for i := 0; i < cntSteeringMembers; i++ {
		vrfList.VrfCheckList[i] = &abci.VrfCheck{
			SteeringMemberCandidateAddress: sdk.ConsAddress(PKS[cntStandingMembers+i].Address()),
			IsVrfTransmission:              true,
		}
	}

	app.DistrKeeper.AllocateTokens(ctx, 30, 30, sdk.ConsAddress(PKS[0].Address()), votes, vrfList)

	for i := 0; i < cntValidators; i++ {
		require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards.IsValid())
	}

	//standing
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3128078817733990130, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards)
	//standing
	for i := 1; i < cntStandingMembers; i++ {
		require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3128078817733990130, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards)
	}
	//steering
	for i := cntStandingMembers; i < cntValidators; i++ {
		require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(3747126436781609170, 16)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[i]).Rewards)
	}

	// end block to bond validator and start new block
	//staking.EndBlocker(ctx, app.StakingKeeper)
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)

	//fmt.Println("total Rewards=> ", app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[1]).Rewards)

	// end period
	endingPeriod := app.DistrKeeper.IncrementValidatorPeriod(ctx, val)

	// calculate delegation rewards ==> 9.091%
	rewards := app.DistrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)
	//fmt.Println("Rewards => ", rewards)
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(255933721450924000, 16)}}, rewards)
}
