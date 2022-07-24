package blogc_test

import (
	"testing"

	keepertest "github.com/andreaZka/blogc/testutil/keeper"
	"github.com/andreaZka/blogc/testutil/nullify"
	"github.com/andreaZka/blogc/x/blogc"
	"github.com/andreaZka/blogc/x/blogc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BlogcKeeper(t)
	blogc.InitGenesis(ctx, *k, genesisState)
	got := blogc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
