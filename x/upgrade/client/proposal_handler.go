package client

import (
	govclient "github.com/reapchain/cosmos-sdk/x/gov/client"
	"github.com/reapchain/cosmos-sdk/x/upgrade/client/cli"
	"github.com/reapchain/cosmos-sdk/x/upgrade/client/rest"
)

var (
	ProposalHandler       = govclient.NewProposalHandler(cli.NewCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
	CancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitCancelUpgradeProposal, rest.ProposalCancelRESTHandler)
)
