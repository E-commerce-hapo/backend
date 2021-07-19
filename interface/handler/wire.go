package handler

import (
	"github.com/google/wire"
	"github.com/kiem-toan/interface/handler/category"
)


var WireSet = wire.NewSet(
	category.New,
)