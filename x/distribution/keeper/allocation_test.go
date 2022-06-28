package keeper_test

import (
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

	app.DistrKeeper.AllocateTokens(ctx, 200, 200, valConsAddr2, votes, vrfList)

	//
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(230, 1)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards)
	//proposer reward + staking.proportional for second proposer = (5 % + )
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(230, 1)}}, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[1]).Rewards)
	//
	require.Equal(t, sdk.DecCoins(sdk.DecCoins(nil)), app.DistrKeeper.GetFeePool(ctx).CommunityPool)
	// 23 * 0.5
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(115, 1)}}, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission)
	// zero commission for second proposer
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[1]).Commission.IsZero())
	//
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(115, 1)}}, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[0]).Rewards)
	// proposer reward + staking.proportional for second proposer = (5 % + 0.5 * (93%)) * 100 = 51.5
	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDecWithPrec(230, 1)}}, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[1]).Rewards)
}

func TestAllocateTokensTruncation(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addrs := simapp.AddTestAddrs(app, ctx, 3, sdk.NewInt(1234))
	valAddrs := simapp.ConvertAddrsToValAddrs(addrs)
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	// create validator with 10% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(1, 1), sdk.NewDecWithPrec(1, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[0], valConsPk1, sdk.NewInt(110), true, stakingtypes.ValidatorTypeStanding)

	// create second validator with 10% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(1, 1), sdk.NewDecWithPrec(1, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[1], valConsPk2, sdk.NewInt(100), true, stakingtypes.ValidatorTypeSteering)

	// create third validator with 10% commission
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(1, 1), sdk.NewDecWithPrec(1, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[2], valConsPk3, sdk.NewInt(100), true, stakingtypes.ValidatorTypeSteering)

	abciValA := abci.Validator{
		Address: valConsPk1.Address(),
		Power:   11,
	}
	abciValB := abci.Validator{
		Address: valConsPk2.Address(),
		Power:   10,
	}
	abciValС := abci.Validator{
		Address: valConsPk3.Address(),
		Power:   10,
	}

	// assert initial state: zero outstanding rewards, zero community pool, zero commission, zero current rewards
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[1]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[1]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetFeePool(ctx).CommunityPool.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[1]).Commission.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[0]).Rewards.IsZero())
	require.True(t, app.DistrKeeper.GetValidatorCurrentRewards(ctx, valAddrs[1]).Rewards.IsZero())

	// allocate tokens as if both had voted and second was proposer
	fees := sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(634195840)))

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
	}

	var vrfList abci.VrfCheckList
	vrfList.VrfCheckList = make([]*abci.VrfCheck, 0)

	app.DistrKeeper.AllocateTokens(ctx, 31, 31, sdk.ConsAddress(valConsPk2.Address()), votes, vrfList)

	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[0]).Rewards.IsValid())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[1]).Rewards.IsValid())
	require.True(t, app.DistrKeeper.GetValidatorOutstandingRewards(ctx, valAddrs[2]).Rewards.IsValid())
}
