package legacytx_test

import (
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/mycodeku/transtionhelper/codec"
	cryptoAmino "github.com/mycodeku/transtionhelper/crypto/codec"
	"github.com/mycodeku/transtionhelper/testutil/testdata"
	sdk "github.com/mycodeku/transtionhelper/types"
	"github.com/mycodeku/transtionhelper/x/auth/migrations/legacytx"
	"github.com/mycodeku/transtionhelper/x/auth/testutil"
)

func testCodec() *codec.LegacyAmino {
	cdc := codec.NewLegacyAmino()
	sdk.RegisterLegacyAminoCodec(cdc)
	cryptoAmino.RegisterCrypto(cdc)
	cdc.RegisterConcrete(&testdata.TestMsg{}, "cosmos-sdk/Test", nil)
	return cdc
}

func TestStdTxConfig(t *testing.T) {
	cdc := testCodec()
	txGen := legacytx.StdTxConfig{Cdc: cdc}
	suite.Run(t, testutil.NewTxConfigTestSuite(txGen))
}
