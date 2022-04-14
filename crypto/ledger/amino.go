package ledger

import (
	"github.com/reapchain/cosmos-sdk/codec"
	cryptoAmino "github.com/reapchain/cosmos-sdk/crypto/codec"
)

var cdc = codec.NewLegacyAmino()

func init() {
	RegisterAmino(cdc)
	cryptoAmino.RegisterCrypto(cdc)
}

// RegisterAmino registers all go-crypto related types in the given (amino) codec.
func RegisterAmino(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(PrivKeyLedgerSecp256k1{},
		"reapchain/PrivKeyLedgerSecp256k1", nil)
}
