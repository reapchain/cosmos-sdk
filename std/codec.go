package std

import (
	"github.com/reapchain/cosmos-sdk/codec"
	"github.com/reapchain/cosmos-sdk/codec/types"
	cryptocodec "github.com/reapchain/cosmos-sdk/crypto/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
	txtypes "github.com/reapchain/cosmos-sdk/types/tx"
)

// RegisterLegacyAminoCodec registers types with the Amino codec.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptocodec.RegisterCrypto(cdc)
	codec.RegisterEvidences(cdc)
}

// RegisterInterfaces registers Interfaces from sdk/types, vesting, crypto, tx.
func RegisterInterfaces(interfaceRegistry types.InterfaceRegistry) {
	sdk.RegisterInterfaces(interfaceRegistry)
	txtypes.RegisterInterfaces(interfaceRegistry)
	cryptocodec.RegisterInterfaces(interfaceRegistry)
}
