//go:build wireinject
// +build wireinject

package main

import (
	"database/sql"

	"github.com/google/wire"

	_ "github.com/lib/pq"

	"github.com/Ak-Srivastav/grpcgo"
)

func InitializeServer() (*grpcgo.Service, error) {
	wire.Build(grpcgo.NewService, provideDB)
	return &grpcgo.Service{}, nil
}

func provideDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://root:pass@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
