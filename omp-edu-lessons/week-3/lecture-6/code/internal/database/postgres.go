package database

import (
	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
)

func NewPostgres(dsn, driver string) (*sqlx.DB, error) {
	db, err := sqlx.Open(driver, dsn)
	if err != nil {
		log.Error().Err(err).Msgf("failed to create database connection")

		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Error().Err(err).Msgf("failed ping the database")

		return nil, err
	}

	return db, nil
}
