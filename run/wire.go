//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"
)

func InitializeServer() (*Service, error) {
	wire.Build(NewService, provideDB)
	return nil, nil
}

func provideDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://root:pass@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
