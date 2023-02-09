package keeper

import (
	"github.com/mezonhub/mezonhub/x/zstaking/types"
)

var _ types.QueryServer = Keeper{}
