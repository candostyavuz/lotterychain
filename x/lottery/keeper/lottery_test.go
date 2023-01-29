package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	keepertest "lotterychain/testutil/keeper"
	"lotterychain/testutil/nullify"
	"lotterychain/x/lottery/keeper"
	"lotterychain/x/lottery/types"
)

func createTestLottery(keeper *keeper.Keeper, ctx sdk.Context) types.Lottery {
	item := types.Lottery{}
	keeper.SetLottery(ctx, item)
	return item
}

func TestLotteryGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	item := createTestLottery(keeper, ctx)
	rst, found := keeper.GetLottery(ctx)
	require.True(t, found)
	require.Equal(t,
		nullify.Fill(&item),
		nullify.Fill(&rst),
	)
}

func TestLotteryRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	createTestLottery(keeper, ctx)
	keeper.RemoveLottery(ctx)
	_, found := keeper.GetLottery(ctx)
	require.False(t, found)
}
