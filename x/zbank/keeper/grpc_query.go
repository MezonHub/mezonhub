package keeper

import (
	"github.com/mezonhub/mezonhub/x/zbank/types"
)

var _ types.QueryServer = Keeper{}
