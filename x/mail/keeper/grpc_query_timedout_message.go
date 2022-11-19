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

func (k Keeper) TimedoutMessageAll(c context.Context, req *types.QueryAllTimedoutMessageRequest) (*types.QueryAllTimedoutMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var timedoutMessages []types.TimedoutMessage
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	timedoutMessageStore := prefix.NewStore(store, types.KeyPrefix(types.TimedoutMessageKey))

	pageRes, err := query.Paginate(timedoutMessageStore, req.Pagination, func(key []byte, value []byte) error {
		var timedoutMessage types.TimedoutMessage
		if err := k.cdc.Unmarshal(value, &timedoutMessage); err != nil {
			return err
		}

		timedoutMessages = append(timedoutMessages, timedoutMessage)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllTimedoutMessageResponse{TimedoutMessage: timedoutMessages, Pagination: pageRes}, nil
}

func (k Keeper) TimedoutMessage(c context.Context, req *types.QueryGetTimedoutMessageRequest) (*types.QueryGetTimedoutMessageResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)
	timedoutMessage, found := k.GetTimedoutMessage(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryGetTimedoutMessageResponse{TimedoutMessage: timedoutMessage}, nil
}
