package keeper

import (
	"github.com/mezonhub/mezonhub/x/rbank/types"
)

var _ types.QueryServer = Keeper{}
