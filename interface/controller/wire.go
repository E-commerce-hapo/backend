package controller

import (
	"github.com/google/wire"
	"github.com/kiem-toan/interface/controller/category"
)

var WireSet = wire.NewSet(
	category.New,
)
