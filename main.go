package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/kiem-toan/cmd/audit-server/build"
	_interface "github.com/kiem-toan/interface"
	"github.com/subosito/gotenv"

	"github.com/kiem-toan/pkg/config"
	"github.com/kiem-toan/pkg/env"
	"github.com/kiem-toan/pkg/integration/consul"
	log2 "github.com/kiem-toan/pkg/log"
	_ "github.com/lib/pq"
)

func init() {
	gotenv.Load()
	env.SetEnvironment(os.Getenv("ENV"))
}

func main() {
	var cfgCh = make(chan config.Config, 1)
	watcher := consul.RegisterConsulWatcher(cfgCh)
	defer watcher.Stop()

	var s *http.Server
	for {
		cfg := <-cfgCh
		go func() {
			config.SetAppConfig(&cfg)
			go log2.RegisterLogStash()
			log2.SetLoglevel(config.GetAppConfig().Log.Level)
			cfg.Info()

			app, err := build.InitApp(cfg)
			if err != nil {
				panic(err)
			}

			routes := _interface.AllRoutes(app)
			var router = _interface.NewRouter(routes)
			headersOk := handlers.AllowedHeaders([]string{"X-Requested-With"})
			originsOk := handlers.AllowedOrigins([]string{`*`})
			methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})
			if s != nil {
				err = s.Shutdown(context.Background())
				if err != nil {
					log2.Panic(err, nil, nil)
				}
			}
			s = &http.Server{
				Addr:    fmt.Sprintf(":%v", config.Port),
				Handler: handlers.CORS(headersOk, originsOk, methodsOk)(router),
			}
			fmt.Println("APP run port: ", config.Port)
			err = s.ListenAndServe()
			switch err {
			case nil, http.ErrServerClosed:
			default:
				log.Fatal(err, nil, nil)
			}
		}()
	}
}
