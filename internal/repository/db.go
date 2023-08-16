package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PostgresConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(cfg PostgresConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode))
	if err != nil {
		return db, err
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		return db, err
	}
	return db, nil
}
