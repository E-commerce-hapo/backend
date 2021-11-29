package dispatcher

import "github.com/google/wire"

var WireSet = wire.NewSet(
	NewDispatcher,
)
