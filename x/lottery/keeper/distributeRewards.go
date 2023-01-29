package keeper

import (
	"crypto/sha256"
	"lotterychain/x/lottery/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) DistributeRewards(ctx sdk.Context) {
	lottery, _ := k.GetLottery(ctx)

	hash := sha256.Sum256([]byte(lottery.TxDataAll))
	lowest16bits := uint16(hash[0])<<8 + uint16(hash[1])
	winnerId := uint64(lowest16bits) % 10

	winner, _ := k.GetParticipant(ctx, winnerId)

	// Check bets
	if winner.Bet == lottery.CurrentMinBet { // no rewards, lottery total pool is carried over
		// reset lottery, keep all the prize pool
		resetLottery := types.Lottery{
			TxCounter:     0,
			TotalFees:     lottery.TotalFees,
			TotalBets:     lottery.TotalBets,
			CurrentMinBet: sdk.NewCoin("token", sdk.NewInt(9223372036854775807)),
			CurrentMaxBet: sdk.NewCoin("token", sdk.ZeroInt()),
			TxDataAll:     "",
		}
		k.SetLottery(ctx, resetLottery)
	} else if winner.Bet == lottery.CurrentMaxBet { // full rewards
		rewardAmount := lottery.TotalBets.Add(lottery.TotalFees)
		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(winner.Address), sdk.Coins{rewardAmount})
		// reset lottery
		resetLottery := types.Lottery{
			TxCounter:     0,
			TotalFees:     sdk.NewCoin("token", sdk.ZeroInt()),
			TotalBets:     sdk.NewCoin("token", sdk.ZeroInt()),
			CurrentMinBet: sdk.NewCoin("token", sdk.NewInt(9223372036854775807)),
			CurrentMaxBet: sdk.NewCoin("token", sdk.ZeroInt()),
			TxDataAll:     "",
		}
		k.SetLottery(ctx, resetLottery)
	} else { // only bets (fees paid remains in lottery pool)
		rewardAmount := lottery.TotalBets
		k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, sdk.AccAddress(winner.Address), sdk.Coins{rewardAmount})
		// reset lottery, keep fees in the prize pool
		resetLottery := types.Lottery{
			TxCounter:     0,
			TotalFees:     lottery.TotalFees,
			TotalBets:     sdk.NewCoin("token", sdk.ZeroInt()),
			CurrentMinBet: sdk.NewCoin("token", sdk.NewInt(9223372036854775807)),
			CurrentMaxBet: sdk.NewCoin("token", sdk.ZeroInt()),
			TxDataAll:     "",
		}
		k.SetLottery(ctx, resetLottery)
	}

}
