package keeper_test

import (
	"context"
	"testing"

	keepertest "Biot/testutil/keeper"
	"Biot/x/biot/keeper"
	"Biot/x/biot/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BiotKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
