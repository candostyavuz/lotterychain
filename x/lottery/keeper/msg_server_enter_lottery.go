package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"lotterychain/x/lottery/types"
)

// ToDo: min - max bet proper update
// ToDo: revert tx after txCounter hits 10
// ToDo: the chosen block proposer can't have any lottery transactions with itself as a sender

// 1. Check if participant is already registered
//	if yes: update bet amount of the participant.
//		txCounter remains same. fee is not refunded. previous bet is REFUNDED. minBet & maxBet adjusted according to new bet
//		(if the same user has new lottery transactions, then only the last one counts, counter doesn’t increase on substitution.)
//	if no: register user as a lottery participant

const requiredFeeInt int64 = 5_000000 // 5token with 6 decimals
const minBetInt int64 = 1_000000      // 1token with 6 decimals
const maxBetInt int64 = 100_000000    // 100token with 6 decimals

func (k msgServer) EnterLottery(goCtx context.Context, msg *types.MsgEnterLottery) (*types.MsgEnterLotteryResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Fetch lottery
	lottery, isFound := k.GetLottery(ctx)
	if !isFound {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "couldn't fetch the lottery object")
	}

	// Check if the address is already registered
	var isRegistered bool = false
	var registerIndex uint64 = 0
	for i := uint64(1); i <= lottery.TxCounter; i++ {
		participant, isFound := k.GetParticipant(ctx, i)
		if isFound {
			if participant.Address == msg.Creator {
				isRegistered = true
				registerIndex = i
				break
			}
		}
	}

	// Check entry base fee
	fee := sdk.NewInt64Coin("token", requiredFeeInt)
	msgFee := msg.Fee

	if msgFee.IsLT(fee) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInsufficientFee, "not enough fee!")
	}

	// Check bet amount
	minBet := sdk.NewInt64Coin("token", minBetInt)
	maxBet := sdk.NewInt64Coin("token", maxBetInt)
	msgBet := msg.Bet

	if msgBet.IsLT(minBet) || maxBet.IsLT(msgBet) {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "bet is out of bounds")
	}

	// Transfer fee + bet into the lottery pool
	participantAddress, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid address!")
	}
	transferAmount := sdk.Coins{msgBet.Add(fee)} // don't send msg.Fee since it can be larger than required fee amount (5token)
	transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, participantAddress, types.ModuleName, transferAmount)
	if transferErr != nil {
		return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "accont to module transfer error!")
	}

	// New participant entering the lottery case
	if !isRegistered {
		// Update lottery object
		txCounter := lottery.TxCounter + 1
		totalFees := lottery.TotalFees.Add(fee)
		totalBets := lottery.TotalBets.Add(msgBet)
		//
		var currMinBet sdk.Coin
		if msgBet.IsLT(lottery.CurrentMinBet) { // No ternary operator in go...
			currMinBet = msgBet
		} else {
			currMinBet = lottery.CurrentMinBet
		}
		//
		var currMaxBet sdk.Coin
		if msgBet.IsGTE(lottery.CurrentMaxBet) { // No ternary operator in go...
			currMaxBet = msgBet
		} else {
			currMaxBet = lottery.CurrentMaxBet
		}
		updatedLottery := types.Lottery{
			TxCounter:     txCounter,
			TotalFees:     totalFees,
			TotalBets:     totalBets,
			CurrentMinBet: currMinBet,
			CurrentMaxBet: currMaxBet,
			TxDataAll:     "", // TBD
		}
		k.SetLottery(ctx, updatedLottery)

		// Update participant object
		newParticipant := types.Participant{
			Id:      txCounter, // starting from 1
			Address: participantAddress.String(),
			Bet:     msgBet,
			TxData:  "", //TBD
		}
		k.SetParticipant(ctx, newParticipant)
	} else { // if the same user has new lottery transactions, then only the last one counts, counter doesn’t increase on substitution.
		participant, _ := k.GetParticipant(ctx, registerIndex)

		totalFees := lottery.TotalFees.Add(fee) // fees are not refunded

		// refund previous bet
		transferRefundErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, participantAddress, sdk.Coins{participant.Bet})
		if transferRefundErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "cannot refund")
		}

		totalBets := (lottery.TotalBets.Sub(participant.Bet)).Add(msgBet) // previous recorded bet is refunded, new bet is added

		// Update participant object
		newParticipant := types.Participant{
			Id:      participant.Id, // starting from 1
			Address: participant.Address,
			Bet:     msgBet,
			TxData:  "", //TBD
		}
		k.SetParticipant(ctx, newParticipant)

		// minBet & maxBet update
		var currMinBet sdk.Coin
		var currMaxBet sdk.Coin

		currMaxBet = k.UpdateMaxBet(ctx)
		currMinBet = k.UpdateMinBet(ctx)

		// Update lottery
		updatedLottery := types.Lottery{
			TxCounter:     lottery.TxCounter,
			TotalFees:     totalFees,
			TotalBets:     totalBets,
			CurrentMinBet: currMinBet,
			CurrentMaxBet: currMaxBet,
			TxDataAll:     "", // TBD
		}
		k.SetLottery(ctx, updatedLottery)

	}

	return &types.MsgEnterLotteryResponse{}, nil
}
