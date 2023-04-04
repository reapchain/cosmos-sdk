package gov_test

import (
	"testing"

	abcitypes "github.com/reapchain/reapchain-core/abci/types"
	tmproto "github.com/reapchain/reapchain-core/proto/reapchain-core/types"
	"github.com/stretchr/testify/require"

	"github.com/reapchain/cosmos-sdk/simapp"
	authtypes "github.com/reapchain/cosmos-sdk/x/auth/types"
	"github.com/reapchain/cosmos-sdk/x/gov/types"
)

func TestItCreatesModuleAccountOnInitBlock(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	app.InitChain(
		abcitypes.RequestInitChain{
			AppStateBytes: []byte("{}"),
			ChainId:       "test-chain-id",
		},
	)

	acc := app.AccountKeeper.GetAccount(ctx, authtypes.NewModuleAddress(types.ModuleName))
	require.NotNil(t, acc)
}
