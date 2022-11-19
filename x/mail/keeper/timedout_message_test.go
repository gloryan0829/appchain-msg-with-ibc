package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	keepertest "planet/testutil/keeper"
	"planet/testutil/nullify"
	"planet/x/mail/keeper"
	"planet/x/mail/types"
)

func createNTimedoutMessage(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.TimedoutMessage {
	items := make([]types.TimedoutMessage, n)
	for i := range items {
		items[i].Id = keeper.AppendTimedoutMessage(ctx, items[i])
	}
	return items
}

func TestTimedoutMessageGet(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNTimedoutMessage(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetTimedoutMessage(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestTimedoutMessageRemove(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNTimedoutMessage(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveTimedoutMessage(ctx, item.Id)
		_, found := keeper.GetTimedoutMessage(ctx, item.Id)
		require.False(t, found)
	}
}

func TestTimedoutMessageGetAll(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNTimedoutMessage(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllTimedoutMessage(ctx)),
	)
}

func TestTimedoutMessageCount(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNTimedoutMessage(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetTimedoutMessageCount(ctx))
}
