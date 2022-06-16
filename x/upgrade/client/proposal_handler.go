package client

import (
	govclient "github.com/mycodeku/transtionhelper/x/gov/client"
	"github.com/mycodeku/transtionhelper/x/upgrade/client/cli"
)

var LegacyProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitLegacyUpgradeProposal)
var LegacyCancelProposalHandler = govclient.NewProposalHandler(cli.NewCmdSubmitLegacyCancelUpgradeProposal)
