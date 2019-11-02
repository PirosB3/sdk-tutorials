package coinswap

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/sdk-tutorials/nameservice/x/coinswap/internal/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

// TODO: ...

// GenesisState - coinswap genesis state
type GenesisState struct {
	Params types.Params `json:"params"`
}

// NewGenesisState is the constructor function for GenesisState
func NewGenesisState(params types.Params) GenesisState {
	return GenesisState{
		Params: params,
	}
}

// DefaultGenesisState creates a default GenesisState object
func DefaultGenesisState() GenesisState {
	return NewGenesisState(types.DefaultParams())
}

// InitGenesis new coinswap genesis
func InitGenesis(ctx sdk.Context, keeper Keeper, data GenesisState) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

// ExportGenesis returns a GenesisState for a given context and keeper.
func ExportGenesis(ctx sdk.Context, keeper Keeper) GenesisState {
	return NewGenesisState(types.DefaultParams())
}

// ValidateGenesis - placeholder function
func ValidateGenesis(data GenesisState) error {
	if err := types.ValidateParams(data.Params); err != nil {
		return err
	}
	return nil
}
