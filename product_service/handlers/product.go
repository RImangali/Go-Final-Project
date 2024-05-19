package handlers

import (
    "context"
    "database/sql"
    "log"
    "github.com/RImangali/product_service/pb"
    _ "github.com/lib/pq"
)

type ProductServer struct {
    pb.UnimplementedProductServiceServer
}

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("localhost", "user=postgres password=Imangali2004 dbname=postgres sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
    query := "INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id"
    var id int
    err := db.QueryRow(query, req.Name, req.Description, req.Price).Scan(&id)
    if err != nil {
        return nil, err
    }
    return &pb.CreateProductResponse{Id: string(id)}, nil
}

func (s *ProductServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
    query := "SELECT id, name, description, price FROM products WHERE id = $1"
    var id int
    var name, description string
    var price float64
    err := db.QueryRow(query, req.Id).Scan(&id, &name, &description, &price)
    if err != nil {
        return nil, err
    }
    return &pb.GetProductResponse{Id: string(id), Name: name, Description: description, Price: price}, nil
}

func (s *ProductServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
    query := "UPDATE products SET name = $1, description = $2, price = $3 WHERE id = $4 RETURNING id"
    var id int
    err := db.QueryRow(query, req.Name, req.Description, req.Price, req.Id).Scan(&id)
    if err != nil {
        return nil, err
    }
    return &pb.UpdateProductResponse{Id: string(id)}, nil
}

func (s *ProductServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
    query := "DELETE FROM products WHERE id = $1 RETURNING id"
    var id int
    err := db.QueryRow(query, req.Id).Scan(&id)
    if err != nil {
        return nil, err
    }
    return &pb.DeleteProductResponse{Id: string(id)}, nil
}
