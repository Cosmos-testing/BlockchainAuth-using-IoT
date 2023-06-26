package keeper

import (
    "context"
    "fmt"

    sdk "github.com/cosmos/cosmos-sdk/types"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"

    "hello/x/hello/types"
)

func (k Keeper) SayCosmos(goCtx context.Context, req *types.QuerySayHelloRequest) (*types.QuerySayHelloResponse, error) {
    if req == nil {
        return nil, status.Error(codes.InvalidArgument, "invalid request")
    }
    ctx := sdk.UnwrapSDKContext(goCtx)
    // TODO: Process the query
    _ = ctx
    return &types.QuerySayCosmosResponse{Name: fmt.Sprintf("hello %s", req.Name)}, nil
}