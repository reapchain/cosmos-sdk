package keeper_test

import (
	"github.com/reapchain/cosmos-sdk/simapp"
	sdk "github.com/reapchain/cosmos-sdk/types"
	authtypes "github.com/reapchain/cosmos-sdk/x/auth/types"
	"github.com/reapchain/cosmos-sdk/x/distribution/types"
)

var (
	PKS = simapp.CreateTestPubKeys(29)

	valConsPk1 = PKS[0]
	valConsPk2 = PKS[1]
	valConsPk3 = PKS[2]
	valConsPk4 = PKS[3]
	valConsPk5 = PKS[4]

	valConsAddr1 = sdk.ConsAddress(valConsPk1.Address())
	valConsAddr2 = sdk.ConsAddress(valConsPk2.Address())
	valConsAddr3 = sdk.ConsAddress(valConsPk3.Address())
	valConsAddr4 = sdk.ConsAddress(valConsPk4.Address())
	valConsAddr5 = sdk.ConsAddress(valConsPk5.Address())

	distrAcc = authtypes.NewEmptyModuleAccount(types.ModuleName)
)
