package keeper

import (
	"lotterychain/x/lottery/types"
)

var _ types.QueryServer = Keeper{}
