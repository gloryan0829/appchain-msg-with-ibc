package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"planet/x/mail/types"
)

// GetTimedoutMessageCount get the total number of timedoutMessage
func (k Keeper) GetTimedoutMessageCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimedoutMessageCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetTimedoutMessageCount set the total number of timedoutMessage
func (k Keeper) SetTimedoutMessageCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.TimedoutMessageCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendTimedoutMessage appends a timedoutMessage in the store with a new id and update the count
func (k Keeper) AppendTimedoutMessage(
	ctx sdk.Context,
	timedoutMessage types.TimedoutMessage,
) uint64 {
	// Create the timedoutMessage
	count := k.GetTimedoutMessageCount(ctx)

	// Set the ID of the appended value
	timedoutMessage.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutMessageKey))
	appendedValue := k.cdc.MustMarshal(&timedoutMessage)
	store.Set(GetTimedoutMessageIDBytes(timedoutMessage.Id), appendedValue)

	// Update timedoutMessage count
	k.SetTimedoutMessageCount(ctx, count+1)

	return count
}

// SetTimedoutMessage set a specific timedoutMessage in the store
func (k Keeper) SetTimedoutMessage(ctx sdk.Context, timedoutMessage types.TimedoutMessage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutMessageKey))
	b := k.cdc.MustMarshal(&timedoutMessage)
	store.Set(GetTimedoutMessageIDBytes(timedoutMessage.Id), b)
}

// GetTimedoutMessage returns a timedoutMessage from its id
func (k Keeper) GetTimedoutMessage(ctx sdk.Context, id uint64) (val types.TimedoutMessage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutMessageKey))
	b := store.Get(GetTimedoutMessageIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveTimedoutMessage removes a timedoutMessage from the store
func (k Keeper) RemoveTimedoutMessage(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutMessageKey))
	store.Delete(GetTimedoutMessageIDBytes(id))
}

// GetAllTimedoutMessage returns all timedoutMessage
func (k Keeper) GetAllTimedoutMessage(ctx sdk.Context) (list []types.TimedoutMessage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.TimedoutMessageKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.TimedoutMessage
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetTimedoutMessageIDBytes returns the byte representation of the ID
func GetTimedoutMessageIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetTimedoutMessageIDFromBytes returns ID in uint64 format from a byte array
func GetTimedoutMessageIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
