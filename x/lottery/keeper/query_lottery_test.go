package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "lotterychain/testutil/keeper"
	"lotterychain/testutil/nullify"
	"lotterychain/x/lottery/types"
)

func TestLotteryQuery(t *testing.T) {
	keeper, ctx := keepertest.LotteryKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	item := createTestLottery(keeper, ctx)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetLotteryRequest
		response *types.QueryGetLotteryResponse
		err      error
	}{
		{
			desc:     "First",
			request:  &types.QueryGetLotteryRequest{},
			response: &types.QueryGetLotteryResponse{Lottery: item},
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Lottery(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}
