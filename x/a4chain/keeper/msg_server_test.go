package keeper_test

import (
	"context"
	"testing"

	keepertest "a4chain/testutil/keeper"
	"a4chain/x/a4chain/keeper"
	"a4chain/x/a4chain/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.A4chainKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
