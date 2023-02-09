package keeper

import (
	"github.com/mezonhub/mezonhub/x/mining/types"
)

var _ types.QueryServer = Keeper{}
