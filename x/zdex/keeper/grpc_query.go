package keeper

import (
	"github.com/mezonhub/mezonhub/x/zdex/types"
)

var _ types.QueryServer = Keeper{}
