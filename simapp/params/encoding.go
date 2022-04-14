package params

import (
	"github.com/reapchain/cosmos-sdk/client"
	"github.com/reapchain/cosmos-sdk/codec"
	"github.com/reapchain/cosmos-sdk/codec/types"
)

// EncodingConfig specifies the concrete encoding types to use for a given app.
// This is provided for compatibility between protobuf and amino implementations.
type EncodingConfig struct {
	InterfaceRegistry types.InterfaceRegistry
	// NOTE: this field will be renamed to Codec
	Marshaler codec.Codec
	TxConfig  client.TxConfig
	Amino     *codec.LegacyAmino
}
