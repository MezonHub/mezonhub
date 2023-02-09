package keeper

import (
	"github.com/mezonhub/mezonhub/x/rdex/types"
)

var _ types.QueryServer = Keeper{}
