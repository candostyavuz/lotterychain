package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"lotterychain/x/lottery/types"
)

// GetParticipantCount get the total number of participant
func (k Keeper) GetParticipantCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ParticipantCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetParticipantCount set the total number of participant
func (k Keeper) SetParticipantCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.ParticipantCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendParticipant appends a participant in the store with a new id and update the count
func (k Keeper) AppendParticipant(
	ctx sdk.Context,
	participant types.Participant,
) uint64 {
	// Create the participant
	count := k.GetParticipantCount(ctx)

	// Set the ID of the appended value
	participant.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParticipantKey))
	appendedValue := k.cdc.MustMarshal(&participant)
	store.Set(GetParticipantIDBytes(participant.Id), appendedValue)

	// Update participant count
	k.SetParticipantCount(ctx, count+1)

	return count
}

// SetParticipant set a specific participant in the store
func (k Keeper) SetParticipant(ctx sdk.Context, participant types.Participant) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParticipantKey))
	b := k.cdc.MustMarshal(&participant)
	store.Set(GetParticipantIDBytes(participant.Id), b)
}

// GetParticipant returns a participant from its id
func (k Keeper) GetParticipant(ctx sdk.Context, id uint64) (val types.Participant, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParticipantKey))
	b := store.Get(GetParticipantIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveParticipant removes a participant from the store
func (k Keeper) RemoveParticipant(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParticipantKey))
	store.Delete(GetParticipantIDBytes(id))
}

// GetAllParticipant returns all participant
func (k Keeper) GetAllParticipant(ctx sdk.Context) (list []types.Participant) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.ParticipantKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Participant
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetParticipantIDBytes returns the byte representation of the ID
func GetParticipantIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetParticipantIDFromBytes returns ID in uint64 format from a byte array
func GetParticipantIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
