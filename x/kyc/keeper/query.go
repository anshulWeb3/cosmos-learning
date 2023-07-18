package keeper

import (
	"example/x/kyc/types"
)

var _ types.QueryServer = Keeper{}
