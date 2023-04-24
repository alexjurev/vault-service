package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/heptiolabs/healthcheck"
	"github.com/patrickmn/go-cache"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.uber.org/zap"

	"github.com/alexjurev/vault-service/config"
	"github.com/alexjurev/vault-service/pkg/infrastructure/restapi"
	"github.com/alexjurev/vault-service/pkg/infrastructure/restapi/endpoints"
	"github.com/alexjurev/vault-service/pkg/logging"
)

func main() {
	// конфиг
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	// логгер
	logger, err := logging.NewLogger(1)
	if err != nil {
		log.Fatal(err)
	}
	// кэш
	cache := cache.New(5*time.Minute, 10*time.Minute)

	// мейн сервер
	errc := make(chan error)
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errc <- fmt.Errorf("%s", <-c)
	}()
	handlers, err := endpoints.NewRespAPIEndpoints(cache, logger)
	if err != nil {
		log.Fatal(err)
	}
	api, err := restapi.NewAPI(
		handlers,
		logger,
		cfg.Transport,
	)
	if err != nil {
		log.Fatal(err)
	}

	// close connections
	api.ServerShutdown = func() {
		_ = logger.Sync()
	}

	go func() {
		logger.Info(fmt.Sprintf("vault server is started on: http://%s:%v", cfg.Transport.Address, cfg.Transport.Port))
		errc <- api.Serve()
	}()
	// дебаг
	healthCheck := healthcheck.NewHandler()
	debugHandler := mux.NewRouter()

	debugHandler.Handle("/metrics", promhttp.Handler()).Methods(http.MethodGet)
	debugHandler.Handle("/live", healthCheck).Methods(http.MethodGet)
	debugHandler.Handle("/ready", healthCheck).Methods(http.MethodGet)

	debugAddr := net.JoinHostPort(cfg.Transport.Address, strconv.Itoa(cfg.Transport.DebugPort))
	logger.Info(fmt.Sprintf("Debug served at http://%s", debugAddr))

	debugServer := http.Server{
		Addr:    debugAddr,
		Handler: debugHandler,
	}
	go func() {
		if err := debugServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Panic("debug listen error", zap.Error(err))
		}
	}()

	logger.With(zap.Error(<-errc)).Error("Exit")
	_ = debugServer.Shutdown(context.Background())
	_ = api.Shutdown()
}
