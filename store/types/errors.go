package types

import (
	sdkerrors "github.com/mycodeku/transtionhelper/types/errors"
)

const StoreCodespace = "store"

var (
	ErrInvalidProof = sdkerrors.Register(StoreCodespace, 2, "invalid proof")
)
