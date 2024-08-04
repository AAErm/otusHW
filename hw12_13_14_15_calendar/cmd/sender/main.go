package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"runtime/debug"
	"syscall"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/config"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/sender"
	sqlstorage "github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage/sql"

	"github.com/imega/mt"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	config := config.NewConfig(configFile)
	if config.Error != nil {
		fmt.Printf("failed to get config with error %v", config.Error)

		return
	}

	logg := logger.New(
		logger.WithLevel(config.Logger.Level),
	)

	defer func() {
		if p := recover(); p != nil {
			logg.Errorf(
				"panic recovered: %s; stack trace: %s",
				p,
				string(debug.Stack()),
			)
		}
	}()

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	storage := sqlstorage.New(
		sqlstorage.WithConfig(config.DB),
		sqlstorage.WithLogger(logg),
	)

	sender := sender.New(
		sender.WithLogger(logg),
		sender.WithStorage(storage),
	)
	mtClient := mt.NewMT(
		mt.WithAMQP(config.AMQP.DSN),
		mt.WithLogger(logg),
		mt.WithConfig(config.AMQP.MtConfig),
	)

	if !HealthCheck(mtClient) {
		logg.Fatalf("failed healthcheck AMQP")
	}

	mtClient.HandleFunc(config.AMQP.ServiceName, sender.Send())

	if err := mtClient.ConnectAndServe(); err != nil {
		logg.Errorf("failed to consume amqp server: %s", err)
	}

	logg.Info("sender is started")
	<-ctx.Done()
	if err := mtClient.Shutdown(); err != nil {
		logg.Error("failed to shutdown mtClient: " + err.Error())
	}
	logg.Info("sender is stopped")
}

var retryCount = 10

func HealthCheck(mtClient mt.MT) bool {
	if !mtClient.HealthCheck() && retryCount > 1 {
		retryCount--
		time.Sleep(1 * time.Second)
		return HealthCheck(mtClient)
	}
	return true
}
