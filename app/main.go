package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/ezzycreative1/svc-url-shortener/app/router"
	"github.com/ezzycreative1/svc-url-shortener/config"
	"github.com/ezzycreative1/svc-url-shortener/pkg/db"
	"github.com/ezzycreative1/svc-url-shortener/pkg/mlog"
)

const version = "1.0.0"
const appName = "url-shortener"

func main() {

	// Load Config ( ENV )
	cfg := config.LoadConfig()

	// Load Logger
	logger := mlog.New("info", "stdout")

	// Initialize Redis
	db.NewRedis(&cfg.Redis)

	// Create Net App Instance
	var (
		host = flag.String("host", "", "host http address to listen on")
		port = flag.String("port", "8081", "port number for http listener")
	)
	flag.Parse()

	addr := net.JoinHostPort(*host, *port)
	srv := http.Server{
		Addr:    addr,
		Handler: router.NewRouter(),
	}

	// Log App Version
	logger.Info("Starting app :" + appName)
	logger.Info("App version  :" + version)

	// Start Server On Goroutine
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Info("shutting down the server")
			panic(fmt.Sprintf("server startup panic: %s", err))
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	// gracefull shutdown stage ===============================================

	logger.Info("shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	// cleanup app ...
	logger.Info("Running cleanup tasks...")

	logger.Info("Done cleanup tasks...")
	logger.Sync()
}
