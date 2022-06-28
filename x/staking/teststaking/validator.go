package teststaking

import (
	"testing"

	"github.com/stretchr/testify/require"

	cryptotypes "github.com/reapchain/cosmos-sdk/crypto/types"
	sdk "github.com/reapchain/cosmos-sdk/types"
	"github.com/reapchain/cosmos-sdk/x/staking/types"
)

// NewValidator is a testing helper method to create validators in tests
func NewValidator(t testing.TB, operator sdk.ValAddress, pubKey cryptotypes.PubKey, valType string) types.Validator {
	v, err := types.NewValidator(operator, pubKey, types.Description{})
	v.Type = valType
	require.NoError(t, err)
	return v
}
