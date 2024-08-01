package main

import (
	"context"
	"fmt"
	"time"

	"github.com/AAErm/otusHW/hw12_13_14_15_calendar/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	retryCount = 20
)

func db_conn(conf config.SqlConf) (*pgxpool.Pool, error) {
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d", conf.User, conf.Password, conf.Host, conf.Port)
	poolConfig, err := pgxpool.ParseConfig(connString)
	if err != nil {
		return nil, err
	}
	db, err := pgxpool.NewWithConfig(context.Background(), poolConfig)
	if err != nil {
		return nil, err
	}
	if err := dbPing(db); err != nil {
		return nil, fmt.Errorf("Unable to connect to database: %s", err)
	}

	return db, nil
}

func dbPing(db *pgxpool.Pool) error {
	err := db.Ping(context.Background())

	if err != nil && retryCount > 1 {
		retryCount--
		time.Sleep(1 * time.Second)
		return dbPing(db)
	}
	if err != nil {
		return err
	}

	return nil
}
