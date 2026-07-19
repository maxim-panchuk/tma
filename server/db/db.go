package db

import (
	"context"
	"errors"
	appconfig "github.com/TON-Market/tma/server/config"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"sync"
	"time"
)

var (
	ErrOpenTransaction   = errors.New("open transaction failed")
	ErrTransactionFailed = errors.New("transaction failed")
	ErrCommitTransaction = errors.New("transaction commit failed")
)

var (
	once      sync.Once
	singleton *pgxpool.Pool
)

func Get() *pgxpool.Pool {
	once.Do(func() {
		config, err := pgxpool.ParseConfig(appconfig.Config.DatabaseURL)
		if err != nil {
			log.Fatalf("parse db config failed: %v", err)
		}

		config.MaxConns = 100
		config.MinConns = 20
		config.MaxConnLifetime = time.Hour
		config.MaxConnIdleTime = 5 * time.Minute
		config.HealthCheckPeriod = 1 * time.Minute

		pool, err := pgxpool.NewWithConfig(context.Background(), config)
		if err != nil {
			log.Fatalf("pool connection create failed: %v", err)
		}

		if _, err = pool.Exec(context.Background(), q); err != nil {
			log.Fatalln(err)
		}

		singleton = pool
	})

	return singleton
}
