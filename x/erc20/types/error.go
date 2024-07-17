package types

import (
    sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
    ErrTokenPairNotFound = sdkerrors.Register(ModuleName, 1, "token pair not found")
)
