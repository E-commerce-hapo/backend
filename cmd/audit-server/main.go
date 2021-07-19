package main

import (
	"fmt"
	"github.com/kiem-toan/cmd/audit-server/build"
	"github.com/kiem-toan/cmd/audit-server/config"
	_interface "github.com/kiem-toan/interface"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

func main() {
	config.InitFlags()
	config.ParseFlags()
	cfg, err := config.Load()
	if err != nil {
		fmt.Errorf("Error in loading config: ", err)
	}
	app, err := build.InitApp(cfg)
	if err != nil {
		panic(err)
	}

	routes := _interface.AllRoutes(app)
	var router = _interface.NewRouter(routes)
	s := &http.Server{
		Addr:    fmt.Sprintf(":%v", cfg.Port),
		Handler: router,
	}

	if err = s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
