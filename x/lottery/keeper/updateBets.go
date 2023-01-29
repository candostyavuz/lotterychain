package keeper

import (
	// "lotterychain/x/lottery/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) UpdateMinBet(ctx sdk.Context) sdk.Coin {
	// Fetch lottery
	lottery, _ := k.GetLottery(ctx)

	firstParticipant, _ := k.GetParticipant(ctx, 1) // updated bets are available in this participant
	smallestBet := firstParticipant.Bet
	for i := uint64(1); i <= lottery.TxCounter; i++ {
		participant, _ := k.GetParticipant(ctx, i)
		if participant.Bet.IsLT(smallestBet) {
			smallestBet = participant.Bet
		}
	}
	return smallestBet
}

func (k Keeper) UpdateMaxBet(ctx sdk.Context) sdk.Coin {
	// Fetch lottery
	lottery, _ := k.GetLottery(ctx)

	firstParticipant, _ := k.GetParticipant(ctx, 1) // updated bets are available in this participant
	biggestBet := firstParticipant.Bet
	for i := uint64(1); i <= lottery.TxCounter; i++ {
		participant, _ := k.GetParticipant(ctx, i)
		if biggestBet.IsLT(participant.Bet) {
			biggestBet = participant.Bet
		}
	}

	return biggestBet
}
