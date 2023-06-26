package biot_test

import (
	"testing"

	keepertest "Biot/testutil/keeper"
	"Biot/testutil/nullify"
	"Biot/x/biot"
	"Biot/x/biot/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BiotKeeper(t)
	biot.InitGenesis(ctx, *k, genesisState)
	got := biot.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
