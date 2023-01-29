package lottery

import (
	"lotterychain/x/lottery/keeper"
	"lotterychain/x/lottery/types"
	"time"

	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

func EndBlocker(ctx sdk.Context, req abci.RequestEndBlock, k keeper.Keeper) {
	defer telemetry.ModuleMeasureSince(types.ModuleName, time.Now(), telemetry.MetricKeyEndBlocker)

	lottery, _ := k.GetLottery(ctx)
	txCounter := lottery.TxCounter

	if lottery.TxCounter >= 10 {
		k.DistributeRewards(ctx)
		// Reset Participants
		for i := uint64(0); i < txCounter; i++ {
			k.RemoveParticipant(ctx, i)
		}
	}

}
