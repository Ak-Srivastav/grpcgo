// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"github.com/Ak-Srivastav/grpcgo"
)

import (
	_ "github.com/lib/pq"
)

// Injectors from wire.go:

func InitializeServer() (*grpcgo.Service, error) {
	db, err := provideDB()
	if err != nil {
		return nil, err
	}
	service := grpcgo.NewService(db)
	return service, nil
}

// wire.go:

func provideDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", "postgresql://root:pass@localhost:5432/postgres?sslmode=disable")
	if err != nil {
		return nil, err
	}
	return db, nil
}
