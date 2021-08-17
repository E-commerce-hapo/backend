package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

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
	// make a new channel to notify on os interrupt of server (ctrl + C)
	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)
	// This blocks the code until the channel receives some message
	sig := <-sigChan
	fmt.Println("Received terminate, graceful shutdown", sig)
	// Once message is consumed shut everything down
	// Gracefully shuts down all client requests. Makes server more reliable
	tc, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	s.Shutdown(tc)
}
