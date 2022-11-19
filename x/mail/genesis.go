package mail

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"planet/x/mail/keeper"
	"planet/x/mail/types"
)

// InitGenesis initializes the module's state from a provided genesis state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the message
	for _, elem := range genState.MessageList {
		k.SetMessage(ctx, elem)
	}

	// Set message count
	k.SetMessageCount(ctx, genState.MessageCount)
	// Set all the sentMessage
	for _, elem := range genState.SentMessageList {
		k.SetSentMessage(ctx, elem)
	}

	// Set sentMessage count
	k.SetSentMessageCount(ctx, genState.SentMessageCount)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetPort(ctx, genState.PortId)
	// Only try to bind to port if it is not already bound, since we may already own
	// port capability from capability InitGenesis
	if !k.IsBound(ctx, genState.PortId) {
		// module binds to the port on InitChain
		// and claims the returned capability
		err := k.BindPort(ctx, genState.PortId)
		if err != nil {
			panic("could not claim port capability: " + err.Error())
		}
	}
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the module's exported genesis
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	genesis.PortId = k.GetPort(ctx)
	genesis.MessageList = k.GetAllMessage(ctx)
	genesis.MessageCount = k.GetMessageCount(ctx)
	genesis.SentMessageList = k.GetAllSentMessage(ctx)
	genesis.SentMessageCount = k.GetSentMessageCount(ctx)
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
