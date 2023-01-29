package lottery_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "lotterychain/testutil/keeper"
	"lotterychain/testutil/nullify"
	"lotterychain/x/lottery"
	"lotterychain/x/lottery/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.LotteryKeeper(t)
	lottery.InitGenesis(ctx, *k, genesisState)
	got := lottery.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
