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

	"github.com/kiem-toan/infrastructure/cmenv"

	"github.com/kiem-toan/cmd/audit-server/build"
	"github.com/kiem-toan/cmd/audit-server/config"
	_interface "github.com/kiem-toan/interface"
	_ "github.com/lib/pq"
)

func main() {
	config.InitFlags()
	config.ParseFlags()
	cfg, err := config.Load()
	if err != nil {
		fmt.Errorf("Error in loading config: ", err)
	}
	cmenv.SetEnvironment("backend-server", cfg.Env)
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

	// Cho ứng dụng của chúng ta chạy background trong 1 Goroutine
	go func() {
		if err = s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Gracefully shutdown
	shutdownGracefully(s)

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
