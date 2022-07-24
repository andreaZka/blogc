package keeper_test

import (
	"testing"

	testkeeper "github.com/andreaZka/blogc/testutil/keeper"
	"github.com/andreaZka/blogc/x/blogc/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BlogcKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
