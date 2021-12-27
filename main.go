package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/kiem-toan/core/route"

	"github.com/gorilla/handlers"
	"github.com/kiem-toan/cmd/audit-server/build"
	"github.com/subosito/gotenv"

	"github.com/kiem-toan/core/config"
	"github.com/kiem-toan/pkg/env"
	log2 "github.com/kiem-toan/pkg/log"
	_ "github.com/lib/pq"
)

func init() {
	gotenv.Load()
	env.SetEnvironment(os.Getenv("ENV"))
	log2.RegisterLogStash(os.Getenv("LOGSTASH_IP"), os.Getenv("LOGSTASH_PORT"), os.Getenv("APPLICATION_NAME"))
}

// @title          Swagger E-Commerce HAPO GoLang API
// @version        1.0
// @description    Bộ API của sàn thương mại điện tử
// @contact.name   Nguyễn Văn Kim Hải
// @contact.url    https://mafc.com.vn/
// @contact.email  kimhai.ngvan@gmail.com
// @license.name   HAPO
// @host           localhost:8080
func main() {
	var s *http.Server
	for {
		bytes, err := config.Asset("config.json")
		if err != nil {
			log2.Fatal(err, nil, nil)
		}
		var cfg config.Config
		err = json.Unmarshal(bytes, &cfg)
		if err != nil {
			log2.Fatal(err, nil, nil)
		}
		cfg.ProjectDir, _ = os.Getwd()
		config.SetAppConfig(cfg)
		app, err := build.InitApp(cfg)
		if err != nil {
			panic(err)
		}

		routes := route.AllRoutes(app)
		var router = route.NewRouter(routes)
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
			Addr:    fmt.Sprintf(":%v", config.ServerPort),
			Handler: handlers.CORS(headersOk, originsOk, methodsOk)(router),
		}
		fmt.Println("APP run port: ", config.ServerPort)
		err = s.ListenAndServe()
		switch err {
		case nil, http.ErrServerClosed:
		default:
			log.Fatal(err, nil, nil)
		}
		// Gracefully shutdown
		shutdownGracefully(s)
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
