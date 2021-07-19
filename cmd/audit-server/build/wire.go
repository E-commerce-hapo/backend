// +build wireinject

package build

import (
	"github.com/google/wire"
	categorying "github.com/kiem-toan/application/category"
	"github.com/kiem-toan/cmd/audit-server/config"
	"github.com/kiem-toan/infrastructure/database"
	_all_controller "github.com/kiem-toan/interface/controller"
	"github.com/kiem-toan/interface/controller/category"
	_all_handler "github.com/kiem-toan/interface/handler"
	category_handler "github.com/kiem-toan/interface/handler/category"
)

func InitApp(cfg config.Config) (*App, error) {
	wire.Build(
		database.WireSet,
		_all_controller.WireSet,
		_all_handler.WireSet,
		categorying.WireSet,
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}

type App struct {
	Db              *database.Database
	CategoryService *category.CategoryService
	CategoryHandler *category_handler.CategoryHandler
}
