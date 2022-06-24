package a4chain_test

import (
	"testing"

	keepertest "a4chain/testutil/keeper"
	"a4chain/testutil/nullify"
	"a4chain/x/a4chain"
	"a4chain/x/a4chain/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.A4chainKeeper(t)
	a4chain.InitGenesis(ctx, *k, genesisState)
	got := a4chain.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
