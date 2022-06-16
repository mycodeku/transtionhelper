package tx

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/mycodeku/transtionhelper/codec"
	codectypes "github.com/mycodeku/transtionhelper/codec/types"
	"github.com/mycodeku/transtionhelper/std"
	"github.com/mycodeku/transtionhelper/testutil/testdata"
	sdk "github.com/mycodeku/transtionhelper/types"
	"github.com/mycodeku/transtionhelper/x/auth/testutil"
)

func TestGenerator(t *testing.T) {
	interfaceRegistry := codectypes.NewInterfaceRegistry()
	std.RegisterInterfaces(interfaceRegistry)
	interfaceRegistry.RegisterImplementations((*sdk.Msg)(nil), &testdata.TestMsg{})
	protoCodec := codec.NewProtoCodec(interfaceRegistry)
	suite.Run(t, testutil.NewTxConfigTestSuite(NewTxConfig(protoCodec, DefaultSignModes)))
}
