package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "github.com/username/mars/testutil/keeper"
	"github.com/username/mars/x/mars/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.MarsKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
