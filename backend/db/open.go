package db

import "github.com/jmoiron/sqlx"

func open(config Settings) (*sqlx.DB, error) {
	db, err := sqlx.Open(postgres, config.String())
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(config.MaxConnections)
	db.SetMaxIdleConns(config.MaxIdleConnections)
	db.SetConnMaxLifetime(config.MaxConnLifetime)

	return db, nil
}
