package keeper

import (
	"github.com/username/mars/x/mars/types"
)

var _ types.QueryServer = Keeper{}
