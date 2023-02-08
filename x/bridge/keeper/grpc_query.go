package keeper

import (
	"github.com/mezonhub/mezonhub/x/bridge/types"
)

var _ types.QueryServer = Keeper{}
