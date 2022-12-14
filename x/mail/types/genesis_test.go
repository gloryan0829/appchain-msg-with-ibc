package types_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	"planet/x/mail/types"
)

func TestGenesisState_Validate(t *testing.T) {
	for _, tc := range []struct {
		desc     string
		genState *types.GenesisState
		valid    bool
	}{
		{
			desc:     "default is valid",
			genState: types.DefaultGenesis(),
			valid:    true,
		},
		{
			desc: "valid genesis state",
			genState: &types.GenesisState{
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
				// this line is used by starport scaffolding # types/genesis/validField
			},
			valid: true,
		},
		{
			desc: "duplicated message",
			genState: &types.GenesisState{
				MessageList: []types.Message{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid message count",
			genState: &types.GenesisState{
				MessageList: []types.Message{
					{
						Id: 1,
					},
				},
				MessageCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated sentMessage",
			genState: &types.GenesisState{
				SentMessageList: []types.SentMessage{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid sentMessage count",
			genState: &types.GenesisState{
				SentMessageList: []types.SentMessage{
					{
						Id: 1,
					},
				},
				SentMessageCount: 0,
			},
			valid: false,
		},
		{
			desc: "duplicated timedoutMessage",
			genState: &types.GenesisState{
				TimedoutMessageList: []types.TimedoutMessage{
					{
						Id: 0,
					},
					{
						Id: 0,
					},
				},
			},
			valid: false,
		},
		{
			desc: "invalid timedoutMessage count",
			genState: &types.GenesisState{
				TimedoutMessageList: []types.TimedoutMessage{
					{
						Id: 1,
					},
				},
				TimedoutMessageCount: 0,
			},
			valid: false,
		},
		// this line is used by starport scaffolding # types/genesis/testcase
	} {
		t.Run(tc.desc, func(t *testing.T) {
			err := tc.genState.Validate()
			if tc.valid {
				require.NoError(t, err)
			} else {
				require.Error(t, err)
			}
		})
	}
}
