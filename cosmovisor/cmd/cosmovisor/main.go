package main

import (
	"os"

	"github.com/mycodeku/transtionhelper/cosmovisor"
	"github.com/mycodeku/transtionhelper/cosmovisor/cmd/cosmovisor/cmd"
	cverrors "github.com/mycodeku/transtionhelper/cosmovisor/errors"
)

func main() {
	cosmovisor.SetupLogging()
	if err := cmd.RunCosmovisorCommand(os.Args[1:]); err != nil {
		cverrors.LogErrors(cosmovisor.Logger, "", err)
		os.Exit(1)
	}
}
