package types

import sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

var (
	ErrPairNotSupported   = sdkerrors.Register(ModuleName, 1, "pair not supported")
	ErrOvertradingLimit   = sdkerrors.Register(ModuleName, 2, "over trading limit")
	ErrQuoteReserveAtZero = sdkerrors.Register(ModuleName, 3, "quote reserve after at zero")
)
