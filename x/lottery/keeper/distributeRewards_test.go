package keeper_test

import (
	"strconv"
	"testing"

	keepertest "lotterychain/testutil/keeper"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"lotterychain/x/lottery/types"
)

func TestDistributeRewards(t *testing.T) {
	k, ctx := keepertest.LotteryKeeper(t)

	{
		lottery := types.Lottery{
			TxCounter:     10,
			TotalFees:     sdk.NewCoin("token", sdk.NewInt(100)),
			TotalBets:     sdk.NewCoin("token", sdk.NewInt(100)),
			CurrentMinBet: sdk.NewCoin("token", sdk.NewInt(10)),
			CurrentMaxBet: sdk.NewCoin("token", sdk.NewInt(20)),
			TxDataAll:     "test",
		}
		k.SetLottery(ctx, lottery)

		// set clients
		for i := int64(1); i <= 10; i++ {
			bet := i * 1_000000
			client := types.Participant{
				Id:      uint64(i),
				Address: ("client" + strconv.Itoa(int(i))),
				Bet:     sdk.NewCoin("token", sdk.NewInt(bet)),
			}
			k.SetParticipant(ctx, client)
		}
		//

		k.DistributeRewards(ctx)

		resetLottery, _ := k.GetLottery(ctx)
		require.Equal(t, resetLottery.TxCounter, 0)
		require.Equal(t, resetLottery.TotalFees, lottery.TotalFees)
		require.Equal(t, resetLottery.TotalBets, sdk.NewCoin("token", sdk.ZeroInt()))
		require.Equal(t, resetLottery.CurrentMinBet, sdk.NewCoin("token", sdk.NewInt(9223372036854775807)))
		require.Equal(t, resetLottery.CurrentMaxBet, sdk.NewCoin("token", sdk.ZeroInt()))
		require.Equal(t, resetLottery.TxDataAll, "")
	}
}
