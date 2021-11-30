package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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
			// Gracefully shutdown
			shutdownGracefully(s)
		}()
	}
}

func shutdownGracefully(s *http.Server) {
	signChan := make(chan os.Signal, 1)
	// Thiết lập một channel để lắng nghe tín hiệu dừng từ hệ điều hành,
	// ở đây chúng ta lưu ý 2 tín hiệu (signal) là SIGINT và SIGTERM
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
	<-signChan
	// Thiết lập một khoản thời gian (Timeout) để dừng hoàn toàn ứng dụng và đóng tất cả kết nối.
	timeWait := 30 * time.Second
	ctx, cancel := context.WithTimeout(context.Background(), timeWait)
	defer func() {
		log.Println("Close another connection")
		cancel()
	}()

	if err := s.Shutdown(ctx); err == context.DeadlineExceeded {
		log.Print("Halted active connections")
	}
	close(signChan)
}