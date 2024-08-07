package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/app"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/config"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
	internalgrpc "github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/server/grpc"
	internalhttp "github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/server/http"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage"
	memorystorage "github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage/memory"
	sqlstorage "github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/storage/sql"
)

var configFile string

func init() {
	flag.StringVar(&configFile, "config", "./configs/config.json", "Path to configuration file")
}

func main() {
	flag.Parse()

	if flag.Arg(0) == "version" {
		printVersion()
		return
	}

	config := config.NewConfig(configFile)
	if config.Error != nil {
		fmt.Printf("failed to get config with error %v", config.Error)

		return
	}

	logg := logger.New(
		logger.WithLevel(config.Logger.Level),
	)

	var storage storage.Storage

	if config.DB.Use {
		storage = sqlstorage.New(
			sqlstorage.WithConfig(config.DB),
			sqlstorage.WithLogger(logg),
		)
	} else {
		storage = memorystorage.New()
	}

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	err := storage.Connect(ctx)
	if err != nil {
		logg.Fatalf("failed to conntect to storage %v", err)
	}
	defer storage.Close(ctx)

	calendar := app.New(
		app.WithLogger(logg),
		app.WithStorage(storage),
	)

	httpServer := internalhttp.NewServer(
		internalhttp.WithLogger(logg),
		internalhttp.WithApplication(calendar),
		internalhttp.WithStorage(storage),
		internalhttp.WithHost(config.Server.Host),
		internalhttp.WithPort(config.Server.Port),
	)

	grpcServer := internalgrpc.New(
		internalgrpc.WithApplication(calendar),
		internalgrpc.WithHost(config.Grpc.Host),
		internalgrpc.WithLogger(logg),
	)

	wg := sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()

		logg.Info("grpc server is running...")

		if err := grpcServer.Start(); err != nil {
			logg.Error("failed to start grpc server: " + err.Error())
			cancel()
			return
		}

		logg.Info("grpc server is stopped")
	}()

	go func() {
		defer wg.Done()

		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := httpServer.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}

		if err := grpcServer.Stop(); err != nil {
			logg.Error("failed to stop grpc server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := httpServer.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}

	wg.Wait()
}
