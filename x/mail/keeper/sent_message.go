package keeper

import (
	"encoding/binary"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"planet/x/mail/types"
)

// GetSentMessageCount get the total number of sentMessage
func (k Keeper) GetSentMessageCount(ctx sdk.Context) uint64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentMessageCountKey)
	bz := store.Get(byteKey)

	// Count doesn't exist: no element
	if bz == nil {
		return 0
	}

	// Parse bytes
	return binary.BigEndian.Uint64(bz)
}

// SetSentMessageCount set the total number of sentMessage
func (k Keeper) SetSentMessageCount(ctx sdk.Context, count uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{})
	byteKey := types.KeyPrefix(types.SentMessageCountKey)
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, count)
	store.Set(byteKey, bz)
}

// AppendSentMessage appends a sentMessage in the store with a new id and update the count
func (k Keeper) AppendSentMessage(
	ctx sdk.Context,
	sentMessage types.SentMessage,
) uint64 {
	// Create the sentMessage
	count := k.GetSentMessageCount(ctx)

	// Set the ID of the appended value
	sentMessage.Id = count

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentMessageKey))
	appendedValue := k.cdc.MustMarshal(&sentMessage)
	store.Set(GetSentMessageIDBytes(sentMessage.Id), appendedValue)

	// Update sentMessage count
	k.SetSentMessageCount(ctx, count+1)

	return count
}

// SetSentMessage set a specific sentMessage in the store
func (k Keeper) SetSentMessage(ctx sdk.Context, sentMessage types.SentMessage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentMessageKey))
	b := k.cdc.MustMarshal(&sentMessage)
	store.Set(GetSentMessageIDBytes(sentMessage.Id), b)
}

// GetSentMessage returns a sentMessage from its id
func (k Keeper) GetSentMessage(ctx sdk.Context, id uint64) (val types.SentMessage, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentMessageKey))
	b := store.Get(GetSentMessageIDBytes(id))
	if b == nil {
		return val, false
	}
	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveSentMessage removes a sentMessage from the store
func (k Keeper) RemoveSentMessage(ctx sdk.Context, id uint64) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentMessageKey))
	store.Delete(GetSentMessageIDBytes(id))
}

// GetAllSentMessage returns all sentMessage
func (k Keeper) GetAllSentMessage(ctx sdk.Context) (list []types.SentMessage) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.SentMessageKey))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.SentMessage
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

// GetSentMessageIDBytes returns the byte representation of the ID
func GetSentMessageIDBytes(id uint64) []byte {
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, id)
	return bz
}

// GetSentMessageIDFromBytes returns ID in uint64 format from a byte array
func GetSentMessageIDFromBytes(bz []byte) uint64 {
	return binary.BigEndian.Uint64(bz)
}
