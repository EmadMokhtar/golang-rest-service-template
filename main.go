package golang_rest_service_template

import (
	"context"
	"github.com/labstack/echo/v4"
	"golang-rest-service-template/config"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() error {
	echoSvc := echo.New()

	// Load configuration using yaml file
	cfg, err := config.LoadConfig()
	if err != nil {
		return err
	}

	// Echo middlewares
	config.SetupMiddlewares(echoSvc, cfg)

	// Start server
	go func() {
		if err = echoSvc.Start(cfg.FullAddress()); err != nil && err != http.ErrServerClosed {
			echoSvc.Logger.Fatal("shutting down the server")
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with a timeout of 10 seconds.
	// Use a buffered channel to avoid missing signals as recommended for signal.Notify
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := echoSvc.Shutdown(ctx); err != nil {
		echoSvc.Logger.Fatal(err)
	}
	return nil

}
