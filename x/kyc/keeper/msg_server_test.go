package keeper_test

import (
	"context"
	"testing"

	keepertest "example/testutil/keeper"
	"example/x/kyc/keeper"
	"example/x/kyc/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.KycKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
