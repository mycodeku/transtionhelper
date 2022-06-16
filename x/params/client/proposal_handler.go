package client

import (
	govclient "github.com/mycodeku/transtionhelper/x/gov/client"
	"github.com/mycodeku/transtionhelper/x/params/client/cli"
)

// ProposalHandler is the param change proposal handler.
var ProposalHandler = govclient.NewProposalHandler(cli.NewSubmitParamChangeProposalTxCmd)
