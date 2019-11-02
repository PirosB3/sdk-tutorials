package coinswap

import (
	"github.com/cosmos/sdk-tutorials/nameservice/x/coinswap/internal/keeper"
	"github.com/cosmos/sdk-tutorials/nameservice/x/coinswap/internal/types"
)

type (
	Keeper             = keeper.Keeper
	MsgSwapOrder       = types.MsgSwapOrder
	MsgAddLiquidity    = types.MsgAddLiquidity
	MsgRemoveLiquidity = types.MsgRemoveLiquidity
)

var (
	ErrInvalidDeadline  = types.ErrInvalidDeadline
	ErrNotPositive      = types.ErrNotPositive
	ErrConstraintNotMet = types.ErrConstraintNotMet
	ModuleCdc           = types.ModuleCdc
	RegisterCodec       = types.RegisterCodec
	NewQuerier          = keeper.NewQuerier
)

const (
	DefaultCodespace = types.DefaultCodespace
	ModuleName       = types.ModuleName
)

const (
	RouterKey = types.RouterKey
	StoreKey  = types.StoreKey
)
