package main

import (
	"fmt"
	"github.com/ashah360/cngo/router"
	"log"
	"net/http"

	"github.com/ashah360/cngo/config"
	lr "github.com/ashah360/cngo/util/logger"
)

func main() {
	appConf := config.AppConfig()

	logger := lr.New(appConf.Debug)

	appRouter := router.New()

	address := fmt.Sprintf(":%d", appConf.Server.Port)

	logger.Info().Msgf("Starting server %v", address)

	s := &http.Server{
		Addr:         address,
		Handler:      appRouter,
		ReadTimeout:  appConf.Server.TimeoutRead,
		WriteTimeout: appConf.Server.TimeoutWrite,
		IdleTimeout:  appConf.Server.TimeoutIdle,
	}

	if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatal("Server startup failed")
	}
}

func main() {
	appConf := config.AppConfig()
	logger := lr.New(appConf.Debug)
}
