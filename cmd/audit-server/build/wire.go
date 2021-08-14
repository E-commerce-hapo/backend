// +build wireinject

package build

import (
	"github.com/google/wire"
	category_app "github.com/kiem-toan/application/category"
	product_app "github.com/kiem-toan/application/product"
	"github.com/kiem-toan/application/product/pm"
	"github.com/kiem-toan/cmd/audit-server/config"
	"github.com/kiem-toan/infrastructure/database"
	"github.com/kiem-toan/infrastructure/event/dispatcher"
	_all_controller "github.com/kiem-toan/interface/controller"
	"github.com/kiem-toan/interface/controller/category"
	"github.com/kiem-toan/interface/controller/product"
	_all_handler "github.com/kiem-toan/interface/handler"
	category_handler "github.com/kiem-toan/interface/handler/category"
	product_handler "github.com/kiem-toan/interface/handler/product"
)

func InitApp(cfg config.Config) (*App, error) {
	wire.Build(
		database.WireSet,
		_all_controller.WireSet,
		_all_handler.WireSet,
		category_app.WireSet,
		product_app.WireSet,
		dispatcher.WireSet,
		wire.Struct(new(App), "*"),
	)
	return &App{}, nil
}

type App struct {
	Db              *database.Database
	CategoryService *category.CategoryService
	CategoryHandler *category_handler.CategoryHandler
	ProductService  *product.ProductService
	ProductHandler  *product_handler.ProductHandler
	ProductManager  *pm.ProductManager
}
