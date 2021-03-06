package testutil

import (
	"fmt"

	"github.com/mycodeku/transtionhelper/client"
	"github.com/mycodeku/transtionhelper/client/flags"
	"github.com/mycodeku/transtionhelper/testutil"
	clitestutil "github.com/mycodeku/transtionhelper/testutil/cli"
	sdk "github.com/mycodeku/transtionhelper/types"
	stakingcli "github.com/mycodeku/transtionhelper/x/staking/client/cli"
)

var commonArgs = []string{
	fmt.Sprintf("--%s=true", flags.FlagSkipConfirmation),
	fmt.Sprintf("--%s=%s", flags.FlagBroadcastMode, flags.BroadcastBlock),
	fmt.Sprintf("--%s=%s", flags.FlagFees, sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))).String()),
}

// MsgRedelegateExec creates a redelegate message.
func MsgRedelegateExec(clientCtx client.Context, from, src, dst, amount fmt.Stringer,
	extraArgs ...string) (testutil.BufferWriter, error) {

	args := []string{
		src.String(),
		dst.String(),
		amount.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from.String()),
		fmt.Sprintf("--%s=%d", flags.FlagGas, 300000),
	}
	args = append(args, extraArgs...)

	args = append(args, commonArgs...)
	return clitestutil.ExecTestCLICmd(clientCtx, stakingcli.NewRedelegateCmd(), args)
}

// MsgUnbondExec creates a unbond message.
func MsgUnbondExec(clientCtx client.Context, from fmt.Stringer, valAddress,
	amount fmt.Stringer, extraArgs ...string) (testutil.BufferWriter, error) {

	args := []string{
		valAddress.String(),
		amount.String(),
		fmt.Sprintf("--%s=%s", flags.FlagFrom, from.String()),
	}

	args = append(args, commonArgs...)
	args = append(args, extraArgs...)
	return clitestutil.ExecTestCLICmd(clientCtx, stakingcli.NewUnbondCmd(), args)
}
