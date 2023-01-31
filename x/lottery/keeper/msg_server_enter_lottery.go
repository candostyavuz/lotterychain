package keeper

import (
	"context"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"lotterychain/x/lottery/types"
)

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

	// Block proposer check  (Designed for single validator network)
	isProposer, err := k.BlockProposerParticipantCheck(ctx, msg.Creator)
	if isProposer {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "proposer can't participate!")
	}

	// Check if the address is already registered
	var isRegistered bool = false
	var registerIndex uint64 = 0
	for i := uint64(0); i < lottery.TxCounter; i++ {
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

		// participant tx data (blocktime info is also appended for increased pseudo randomness)
		txData := msg.String() + ctx.BlockTime().Format(time.UnixDate)

		// Create participant object
		newParticipant := types.Participant{
			Id:      lottery.TxCounter,
			Address: participantAddress.String(),
			Bet:     msgBet,
			TxData:  txData,
		}
		k.SetParticipant(ctx, newParticipant)

		// Update lottery object
		txCounter := lottery.TxCounter + 1
		updatedLottery := types.Lottery{
			TxCounter:     txCounter,
			TotalFees:     totalFees,
			TotalBets:     totalBets,
			CurrentMinBet: currMinBet,
			CurrentMaxBet: currMaxBet,
			TxDataAll:     lottery.TxDataAll + txData,
			LastWinner:    lottery.LastWinner,
			LastWinnerIdx: lottery.LastWinnerIdx,
		}
		k.SetLottery(ctx, updatedLottery)

	} else { // if the same user has new lottery transactions, then only the last one counts, counter doesn’t increase on substitution.
		participant, _ := k.GetParticipant(ctx, registerIndex)

		totalFees := lottery.TotalFees.Add(fee) // fees are not refunded

		// refund previous bet
		transferRefundErr := k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, participantAddress, sdk.Coins{participant.Bet})
		if transferRefundErr != nil {
			return nil, sdkerrors.Wrap(sdkerrors.ErrPanic, "cannot refund")
		}

		totalBets := (lottery.TotalBets.Sub(participant.Bet)).Add(msgBet) // previous recorded bet is refunded, new bet is added

		// tx Data is appended in the case of multiple tx coming from the same address
		txData := msg.String() + ctx.BlockTime().Format(time.UnixDate)

		// Update participant object
		newParticipant := types.Participant{
			Id:      participant.Id,
			Address: participant.Address,
			Bet:     msgBet,
			TxData:  participant.TxData + txData,
		}
		k.SetParticipant(ctx, newParticipant)

		// minBet & maxBet update
		var currMinBet sdk.Coin
		var currMaxBet sdk.Coin

		currMaxBet = k.UpdateMaxBet(ctx)
		currMinBet = k.UpdateMinBet(ctx)

		// Update lottery object
		updatedLottery := types.Lottery{
			TxCounter:     lottery.TxCounter,
			TotalFees:     totalFees,
			TotalBets:     totalBets,
			CurrentMinBet: currMinBet,
			CurrentMaxBet: currMaxBet,
			TxDataAll:     lottery.TxDataAll + txData,
			LastWinner:    lottery.LastWinner,
			LastWinnerIdx: lottery.LastWinnerIdx,
		}
		k.SetLottery(ctx, updatedLottery)

	}

	return &types.MsgEnterLotteryResponse{}, nil
}
