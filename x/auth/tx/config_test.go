package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/reapchain/cosmos-sdk/codec"
	codectypes "github.com/reapchain/cosmos-sdk/codec/types"
	"github.com/reapchain/cosmos-sdk/std"
	"github.com/reapchain/cosmos-sdk/testutil/testdata"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
