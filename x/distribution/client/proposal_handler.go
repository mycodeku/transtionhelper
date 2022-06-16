package client

import (
	"github.com/mycodeku/transtionhelper/x/distribution/client/cli"
	govclient "github.com/mycodeku/transtionhelper/x/gov/client"
)

// ProposalHandler is the community spend proposal handler.
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal)
)
