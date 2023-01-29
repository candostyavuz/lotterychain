package lottery

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lotterychain/x/lottery/keeper"
	"lotterychain/x/lottery/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the participant
	for _, elem := range genState.ParticipantList {
		k.SetParticipant(ctx, elem)
	}

	// Set participant count
	k.SetParticipantCount(ctx, genState.ParticipantCount)
	// Set if defined
	if genState.Lottery != nil {
		k.SetLottery(ctx, *genState.Lottery)
	}
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.ParticipantList = k.GetAllParticipant(ctx)
	genesis.ParticipantCount = k.GetParticipantCount(ctx)
	// Get all lottery
	lottery, found := k.GetLottery(ctx)
	if found {
		genesis.Lottery = &lottery
	}
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
