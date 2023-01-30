package keeper

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/bech32"
)

func (k Keeper) BlockProposerParticipantCheck(ctx sdk.Context, msgCreator string) (bool, error) {
	validators := k.stakingKeeper.GetLastValidators(ctx)
	valAddr := validators[0].OperatorAddress
	if valAddr == "" {
		return false, sdkerrors.Wrap(sdkerrors.ErrPanic, "validator couldn't be fetched")
	}

	hrp, data, err := bech32.DecodeAndConvert(msgCreator)
	if err != nil {
		return false, sdkerrors.Wrap(sdkerrors.ErrPanic, "decode err")
	}
	// Change the HRP to "cosmosvaloper..."
	hrp = "cosmosvaloper"
	// Encode the data back into a bech32 string
	msgCreatorValAddr, err2 := bech32.ConvertAndEncode(hrp, data)
	if err2 != nil {
		return false, sdkerrors.Wrap(sdkerrors.ErrPanic, "encode err")
	}

	if valAddr == msgCreatorValAddr {
		return true, sdkerrors.Wrap(sdkerrors.ErrPanic, "proposer can't participate!")
	} else {
		return false, nil
	}
}
