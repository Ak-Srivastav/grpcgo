package grpcgo

import (
	"context"
	"database/sql"
	"log"
)

type Service struct {
	UnimplementedCrudServiceServer
	db *sql.DB
}

func NewService(db *sql.DB) *Service {
	return &Service{db: db}
}

// Returning temporary data as of now

func (s *Service) CreateItem(ctx context.Context, req *CreateItemRequest) (*CreateItemResponse, error) {
	// Implement your logic to create an item in CockroachDB
	log.Printf("Received CreateItem request: %v", req)
	_, err := s.db.ExecContext(ctx, "INSERT INTO items (name, description) VALUES ($1, $2)", req.Name, req.Description)
	if err != nil {
		return nil, err
	}
	return &CreateItemResponse{Id: "1"}, nil
}

func (s *Service) ReadItem(ctx context.Context, req *ReadItemRequest) (*ReadItemResponse, error) {
	// Implement your logic to read an item from CockroachDB
	log.Printf("Received ReadItem request: %v", req)
	var name, description string
	err := s.db.QueryRowContext(ctx, "SELECT name, description FROM items WHERE id = $1", req.Id).Scan(&name, &description)
	if err != nil {
		return nil, err
	}
	return &ReadItemResponse{Id: req.Id, Name: name, Description: description}, nil
}

func (s *Service) UpdateItem(ctx context.Context, req *UpdateItemRequest) (*UpdateItemResponse, error) {
	// Implement your logic to update an item in CockroachDB
	log.Printf("Received UpdateItem request: %v", req)
	_, err := s.db.ExecContext(ctx, "UPDATE items SET name = $1, description = $2 WHERE id = $3", req.Name, req.Description, req.Id)
	if err != nil {
		return nil, err
	}
	return &UpdateItemResponse{Success: true}, nil
}

func (s *Service) DeleteItem(ctx context.Context, req *DeleteItemRequest) (*DeleteItemResponse, error) {
	// Implement your logic to delete an item from CockroachDB
	log.Printf("Received DeleteItem request: %v", req)
	_, err := s.db.ExecContext(ctx, "DELETE FROM items WHERE id = $1", req.Id)
	if err != nil {
		return nil, err
	}
	return &DeleteItemResponse{Success: true}, nil
}
