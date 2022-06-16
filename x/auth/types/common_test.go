package types_test

import (
	"github.com/mycodeku/transtionhelper/simapp"
)

var (
	ecdc                  = simapp.MakeTestEncodingConfig()
	appCodec, legacyAmino = ecdc.Codec, ecdc.Amino
)
