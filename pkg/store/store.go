package store

import (
	"github.com/jmoiron/sqlx"
)

type Config struct {
	Hostname string
	Port     int
	Username string
	Password string
}

type Store struct {
	DB sqlx.DB
}
