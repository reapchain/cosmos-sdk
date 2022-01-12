package testutil

import (
	"github.com/reapchain/cosmos-sdk/testutil"
	clitestutil "github.com/reapchain/cosmos-sdk/testutil/cli"
	"github.com/reapchain/cosmos-sdk/testutil/network"
	"github.com/reapchain/cosmos-sdk/x/authz/client/cli"
)

func ExecGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
