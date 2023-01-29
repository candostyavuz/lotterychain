package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lotterychain/x/lottery/types"
)

// SetLottery set lottery in the store
func (k Keeper) SetLottery(ctx sdk.Context, lottery types.Lottery) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKey))
	b := k.cdc.MustMarshal(&lottery)
	store.Set([]byte{0}, b)
}

// GetLottery returns lottery
func (k Keeper) GetLottery(ctx sdk.Context) (val types.Lottery, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKey))

	b := store.Get([]byte{0})
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveLottery removes lottery from the store
func (k Keeper) RemoveLottery(ctx sdk.Context) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.LotteryKey))
	store.Delete([]byte{0})
}
