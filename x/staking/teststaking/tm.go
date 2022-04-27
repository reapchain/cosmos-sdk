package teststaking

import (
	tmcrypto "github.com/reapchain/reapchain-core/crypto"
	tmtypes "github.com/reapchain/reapchain-core/types"

	cryptocodec "github.com/reapchain/cosmos-sdk/crypto/codec"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/x/staking/types"
)

// GetTmConsPubKey gets the validator's public key as a tmcrypto.PubKey.
func GetTmConsPubKey(v types.Validator) (tmcrypto.PubKey, error) {
	pk, err := v.ConsPubKey()
	if err != nil {
		return nil, err
	}

	return cryptocodec.ToTmPubKeyInterface(pk)
}

// ToTmValidator casts an SDK validator to a tendermint type Validator.
func ToTmValidator(v types.Validator, r sdk.Int) (*tmtypes.Validator, error) {
	tmPk, err := GetTmConsPubKey(v)
	if err != nil {
		return nil, err
	}

	return tmtypes.NewValidator(tmPk, v.ConsensusPower(r), "standing"), nil
}

// ToTmValidators casts all validators to the corresponding tendermint type.
func ToTmValidators(v types.Validators, r sdk.Int) ([]*tmtypes.Validator, error) {
	validators := make([]*tmtypes.Validator, len(v))
	var err error
	for i, val := range v {
		validators[i], err = ToTmValidator(val, r)
		if err != nil {
			return nil, err
		}
	}

	return validators, nil
}
