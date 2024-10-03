package profile

import (
	"github.com/kitty314/1.17.0/common/atomic"
)

// StoreSelected is a global switch for storing selected proxy to cache
var StoreSelected = atomic.NewBool(true)
