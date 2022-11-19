package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"planet/x/mail/types"
)

func (k Keeper) SentMessageAll(c context.Context, req *types.QueryAllSentMessageRequest) (*types.QueryAllSentMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var sentMessages []types.SentMessage
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	sentMessageStore := prefix.NewStore(store, types.KeyPrefix(types.SentMessageKey))

	pageRes, err := query.Paginate(sentMessageStore, req.Pagination, func(key []byte, value []byte) error {
		var sentMessage types.SentMessage
		if err := k.cdc.Unmarshal(value, &sentMessage); err != nil {
			return err
		}

		sentMessages = append(sentMessages, sentMessage)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllSentMessageResponse{SentMessage: sentMessages, Pagination: pageRes}, nil
}

func (k Keeper) SentMessage(c context.Context, req *types.QueryGetSentMessageRequest) (*types.QueryGetSentMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	sentMessage, found := k.GetSentMessage(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetSentMessageResponse{SentMessage: sentMessage}, nil
}
