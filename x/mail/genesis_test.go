package mail_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "planet/testutil/keeper"
	"planet/testutil/nullify"
	"planet/x/mail"
	"planet/x/mail/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		MessageList: []types.Message{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		MessageCount: 2,
		SentMessageList: []types.SentMessage{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		SentMessageCount: 2,
		TimedoutMessageList: []types.TimedoutMessage{
			{
				Id: 0,
			},
			{
				Id: 1,
			},
		},
		TimedoutMessageCount: 2,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.MailKeeper(t)
	mail.InitGenesis(ctx, *k, genesisState)
	got := mail.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.MessageList, got.MessageList)
	require.Equal(t, genesisState.MessageCount, got.MessageCount)
	require.ElementsMatch(t, genesisState.SentMessageList, got.SentMessageList)
	require.Equal(t, genesisState.SentMessageCount, got.SentMessageCount)
	require.ElementsMatch(t, genesisState.TimedoutMessageList, got.TimedoutMessageList)
	require.Equal(t, genesisState.TimedoutMessageCount, got.TimedoutMessageCount)
	// this line is used by starport scaffolding # genesis/test/assert
}
