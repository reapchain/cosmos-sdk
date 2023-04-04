package types

import (
	"github.com/reapchain/cosmos-sdk/codec"
	cryptocodec "github.com/reapchain/cosmos-sdk/crypto/codec"
)

var amino = codec.NewLegacyAmino()

func init() {
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
