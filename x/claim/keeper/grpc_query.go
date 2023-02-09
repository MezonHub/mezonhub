package keeper

import (
	"github.com/mezonhub/mezonhub/x/claim/types"
)

var _ types.QueryServer = Keeper{}
