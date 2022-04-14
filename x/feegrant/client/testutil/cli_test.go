// +build norace

package testutil

import (
	"testing"

	"github.com/reapchain/cosmos-sdk/testutil/network"

	"github.com/stretchr/testify/suite"
)

func TestIntegrationTestSuite(t *testing.T) {
	cfg := network.DefaultConfig()
	cfg.NumValidators = 3
	suite.Run(t, NewIntegrationTestSuite(cfg))
}
