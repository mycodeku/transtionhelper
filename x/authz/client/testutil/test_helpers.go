package testutil

import (
	"github.com/mycodeku/transtionhelper/testutil"
	clitestutil "github.com/mycodeku/transtionhelper/testutil/cli"
	"github.com/mycodeku/transtionhelper/testutil/network"
	"github.com/mycodeku/transtionhelper/x/authz/client/cli"
)

func CreateGrant(val *network.Validator, args []string) (testutil.BufferWriter, error) {
	cmd := cli.NewCmdGrantAuthorization()
	clientCtx := val.ClientCtx
	return clitestutil.ExecTestCLICmd(clientCtx, cmd, args)
}
