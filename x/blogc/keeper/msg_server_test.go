package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/andreaZka/blogc/testutil/keeper"
	"github.com/andreaZka/blogc/x/blogc/keeper"
	"github.com/andreaZka/blogc/x/blogc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BlogcKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
