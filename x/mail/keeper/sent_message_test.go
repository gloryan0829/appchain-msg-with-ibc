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

func createNSentMessage(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.SentMessage {
	items := make([]types.SentMessage, n)
	for i := range items {
		items[i].Id = keeper.AppendSentMessage(ctx, items[i])
	}
	return items
}

func TestSentMessageGet(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNSentMessage(keeper, ctx, 10)
	for _, item := range items {
		got, found := keeper.GetSentMessage(ctx, item.Id)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&got),
		)
	}
}

func TestSentMessageRemove(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNSentMessage(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveSentMessage(ctx, item.Id)
		_, found := keeper.GetSentMessage(ctx, item.Id)
		require.False(t, found)
	}
}

func TestSentMessageGetAll(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNSentMessage(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllSentMessage(ctx)),
	)
}

func TestSentMessageCount(t *testing.T) {
	keeper, ctx := keepertest.MailKeeper(t)
	items := createNSentMessage(keeper, ctx, 10)
	count := uint64(len(items))
	require.Equal(t, count, keeper.GetSentMessageCount(ctx))
}
