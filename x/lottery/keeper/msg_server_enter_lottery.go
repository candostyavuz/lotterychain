package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"lotterychain/x/lottery/types"
)

// Valid only when sender has enough funds to cover lottery fee + minimal bet
// ToDo: revert tx after txCounter hits 10
// ToDo: the chosen block proposer can't have any lottery transactions with itself as a sender

// 1. Check if participant is already registered
//	if yes: update bet amount (if the same user has new lottery transactions, then only the last one counts, counter doesn’t increase on substitution.)
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
	for i := uint64(1); i <= lottery.TxCounter; i++ {
		participant, isFound := k.GetParticipant(ctx, i)
		if isFound {
			if participant.Address == msg.Creator {
				isRegistered = true
				break
			}
		}
	}

	// New participant entering the lottery case
	if !isRegistered {
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
		creatorAddress, err := sdk.AccAddressFromBech32(msg.Creator)
		if err != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "invalid address!")
		}
		transferAmount := sdk.Coins{msgBet.Add(msgFee)}
		transferErr := k.bankKeeper.SendCoinsFromAccountToModule(ctx, creatorAddress, types.ModuleName, transferAmount)
		if transferErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "accont to module transfer error!")
		}

		// Update lottery object
		txCounter := lottery.TxCounter + 1
		totalFees := lottery.TotalFees.Add(msgFee)
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
			Address: creatorAddress.String(),
			Bet:     msgBet,
			TxData:  "", //TBD
		}
		k.SetParticipant(ctx, newParticipant)
	} else {

	}

	return &types.MsgEnterLotteryResponse{}, nil
}
