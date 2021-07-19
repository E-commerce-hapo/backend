// +build wireinject

package build

import (
	"github.com/google/wire"
	"github.com/kiem-toan/cmd/audit-server/config"
	"github.com/kiem-toan/infrastructure/database"
	_all_controller"github.com/kiem-toan/interface/controller"
	_all_handler"github.com/kiem-toan/interface/handler"
	"github.com/kiem-toan/interface/controller/category"
	category_handler"github.com/kiem-toan/interface/handler/category"
)

func InitApp(cfg config.Config) (*App, error) {
	wire.Build(
		database.WireSet,
		_all_controller.WireSet,
		_all_handler.WireSet,
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}

type App struct {
	Db              *database.Database
	CategoryService *category.CategoryService
	CategoryHandler *category_handler.CategoryHandler
}
