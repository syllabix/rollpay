package db

import (
	"context"
	"embed"
	"fmt"
	"net/http"
	"time"

	"github.com/jmoiron/sqlx"
	migrate "github.com/rubenv/sql-migrate"
	"go.uber.org/fx"
	"go.uber.org/zap"

	// postgres driver required import
	_ "github.com/lib/pq"
)

//go:embed .migrations/rollpay/*
var migrations embed.FS

const (
	postgres        = "postgres"
	migrationsTable = "migrations"
)

// Rollpay is the primary database for the application.
type Rollpay struct {
	*sqlx.DB
}

// SetupRollpay attempts configure and connect to the Rollpay database
func SetupRollpay(lc fx.Lifecycle, config Settings, log *zap.Logger) (Rollpay, error) {
	db, err := open(config)
	if err != nil {
		return Rollpay{}, fmt.Errorf("unable to establish a connection with rollpay db: %w", err)
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			err = ensure(db, log)
			if err != nil {
				return fmt.Errorf("unable to establish a connection with rollpay db: %w", err)
			}

			count, err := runMigrations(db)
			if err != nil {
				return fmt.Errorf("database migrations failed: %w", err)
			}
			log.Info("rollpay database migrations completed ok", zap.Int("count", count))
			return nil
		},
		OnStop: func(ctx context.Context) error {
			err := db.Close()
			if err != nil {
				return fmt.Errorf("database connection failed to close properly: %w", err)
			}
			return nil
		},
	})

	return Rollpay{db}, nil
}

func ensure(db *sqlx.DB, log *zap.Logger) (err error) {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	retries := 0
	for retries < 10 {
		err = db.Ping()
		if err != nil {
			log.Warn("unable to establish postgres db connection: retrying...", zap.Error(err))
			<-ticker.C
			retries++
		} else {
			break
		}
	}

	return err
}

func runMigrations(db *sqlx.DB) (count int, err error) {
	source := migrate.HttpFileSystemMigrationSource{
		FileSystem: (http.FS(migrations)),
	}

	migrate.SetTable(migrationsTable)
	count, err = migrate.Exec(db.DB, postgres, source, migrate.Up)
	if err != nil {
		return 0, err
	}

	return count, err
}
