package slashing_test

import (
	sdk "github.com/reapchain/cosmos-sdk/types"
)

// The default power validators are initialized to have within tests
var InitTokens = sdk.TokensFromConsensusPower(200, sdk.DefaultPowerReduction)
