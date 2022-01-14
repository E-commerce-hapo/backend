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

	_ "github.com/lib/pq"
	"github.com/subosito/gotenv"

	"github.com/E-commerce-hapo/backend/core/config"
	"github.com/E-commerce-hapo/backend/pkg/env"
	log2 "github.com/E-commerce-hapo/backend/pkg/log"
	"github.com/E-commerce-hapo/backend/registry"
	server2 "github.com/E-commerce-hapo/backend/server"
)

func init() {
	gotenv.Load()
	env.SetEnvironment(os.Getenv("ENV"))
	//log2.RegisterLogStash(os.Getenv("LOGSTASH_IP"), os.Getenv("LOGSTASH_PORT"), os.Getenv("APPLICATION_NAME"))
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
		//_, err = build.InitApp(cfg)
		if err != nil {
			panic(err)
		}

		r, err := registry.New(cfg)
		if err != nil {
			log2.Fatal(err, nil, nil)
		}
		if err = r.MigratePostgres(); err != nil {
			log2.Error(err, nil, nil)
		}
		server := server2.New(r)

		if s != nil {
			err = s.Shutdown(context.Background())
			if err != nil {
				log2.Panic(err, nil, nil)
			}
		}
		s = &http.Server{
			Addr:    fmt.Sprintf(":%v", config.ServerPort),
			Handler: server,
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
