package kyc_test

import (
	"testing"

	keepertest "example/testutil/keeper"
	"example/testutil/nullify"
	"example/x/kyc"
	"example/x/kyc/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.KycKeeper(t)
	kyc.InitGenesis(ctx, *k, genesisState)
	got := kyc.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
