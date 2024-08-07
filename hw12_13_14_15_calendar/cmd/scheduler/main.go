package main

import (
	"context"
	"flag"
	"fmt"
	"os/signal"
	"runtime/debug"
	"syscall"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/config"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/scheduler"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"

	memorystorage "github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage/sql"
	"github.com/go-co-op/gocron"
	"github.com/imega/mt"

	"time"
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

	var storage storage.Storage

	if config.DB.Use {
		storage = sqlstorage.New(
			sqlstorage.WithConfig(config.DB),
			sqlstorage.WithLogger(logg),
		)
	} else {
		storage = memorystorage.New(
			memorystorage.WithLogger(logg),
		)
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	err := storage.Connect(ctx)
	if err != nil {
		logg.Fatalf("failed to conntect to storage %v", err)
	}
	defer storage.Close(ctx)

	mtClient := mt.NewMT(
		mt.WithAMQP(config.AMQP.DSN),
		mt.WithLogger(logg),
		mt.WithConfig(config.AMQP.MtConfig),
	)

	if !HealthCheck(mtClient) {
		logg.Fatalf("failed healthcheck AMQP")
	}

	scheduler := scheduler.New(
		scheduler.WithLogger(logg),
		scheduler.WithMt(mtClient),
		scheduler.WithStorage(storage),
		scheduler.WithInterval(time.Duration(config.Scheduler.Interval)*time.Second),
		scheduler.WithServiceName(config.AMQP.ServiceName),
	)

	if err := mtClient.ConnectAndServe(); err != nil {
		logg.Errorf("failed to consume amqp server: %s", err)
	}

	s := gocron.NewScheduler(time.UTC)
	_, err = s.Every(config.Scheduler.Interval).Seconds().Do(scheduler.Run())
	if err != nil {
		logg.Fatalf("scheduler is broken %v", err)
	}
	s.StartAsync()

	defer s.Stop()

	logg.Info("scheduler is started")

	<-ctx.Done()
	if err := mtClient.Shutdown(); err != nil {
		logg.Error("failed to shutdown mtClient: " + err.Error())
	}

	logg.Info("service is stopped")
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
