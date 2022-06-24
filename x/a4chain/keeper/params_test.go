package keeper_test

import (
	"testing"

	testkeeper "a4chain/testutil/keeper"
	"a4chain/x/a4chain/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.A4chainKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
