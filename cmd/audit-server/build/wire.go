// +build wireinject

package build

import (
	"github.com/google/wire"
	category_app "github.com/kiem-toan/application/category"
	_all_controller "github.com/kiem-toan/interface/controller"
	"github.com/kiem-toan/interface/controller/category"
	_all_handler "github.com/kiem-toan/interface/handler"
	category_handler "github.com/kiem-toan/interface/handler/category"
	"github.com/kiem-toan/pkg/config"
	"github.com/kiem-toan/pkg/database"
	"github.com/kiem-toan/pkg/event/dispatcher"
	"github.com/kiem-toan/pkg/integration/email"
)

func InitApp(cfg config.Config) (*App, error) {
	wire.Build(
		database.WireSet,
		_all_controller.WireSet,
		_all_handler.WireSet,
		category_app.WireSet,
		dispatcher.WireSet,
		email.WireSet,
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}

type App struct {
	Db              *database.Database
	CategoryService *category.CategoryService
	CategoryHandler *category_handler.CategoryHandler
}
