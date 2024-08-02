package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/app"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/config"
	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/logger"
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

	if config.Sql.Use {
		conn, err := db_conn(config.Sql)
		if err != nil {
			logg.Fatalf("Unable to connect to database: %v", err)
		}
		storage = sqlstorage.New(
			sqlstorage.WithConnect(conn),
			sqlstorage.WithLogger(logg),
		)
	} else {
		storage = memorystorage.New()
	}

	calendar := app.New(
		app.WithLogger(logg),
		app.WithStorage(&storage),
	)

	server := internalhttp.NewServer(
		internalhttp.WithLogger(logg),
		internalhttp.WithApplication(calendar),
	)

	ctx, cancel := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP)
	defer cancel()

	go func() {
		<-ctx.Done()

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
		defer cancel()

		if err := server.Stop(ctx); err != nil {
			logg.Error("failed to stop http server: " + err.Error())
		}
	}()

	logg.Info("calendar is running...")

	if err := server.Start(ctx); err != nil {
		logg.Error("failed to start http server: " + err.Error())
		cancel()
		os.Exit(1) //nolint:gocritic
	}
}
