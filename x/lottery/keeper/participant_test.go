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

func createNParticipant(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Participant {
	items := make([]types.Participant, n)
	for i := range items {
		items[i].Id = keeper.AppendParticipant(ctx, items[i])
	}
	return items
}

func TestParticipantGet(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNParticipant(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetParticipant(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestParticipantRemove(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNParticipant(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveParticipant(ctx, item.Id)
		_, found := keeper.GetParticipant(ctx, item.Id)
		require.False(t, found)
	}
}

func TestParticipantGetAll(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNParticipant(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllParticipant(ctx)),
	)
}

func TestParticipantCount(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	items := createNParticipant(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetParticipantCount(ctx))
}
