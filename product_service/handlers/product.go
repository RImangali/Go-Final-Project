package handlers

import (
    "context"
    "database/sql"
    "log"

    "github.com/RImangali/product_service/pb"
)

type ProductServer struct {
    pb.UnimplementedProductServiceServer
    db *sql.DB
}

func NewProductServer(db *sql.DB) *ProductServer {
    return &ProductServer{db: db}
}

func (s *ProductServer) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
    var id string
    err := s.db.QueryRow("INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id", req.Name, req.Description, req.Price).Scan(&id)
    if err != nil {
        log.Printf("Failed to create product: %v", err)
        return nil, err
    }
    return &pb.CreateProductResponse{Id: id}, nil
}

func (s *ProductServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
    var product pb.GetProductResponse
    err := s.db.QueryRow("SELECT id, name, description, price FROM products WHERE id = $1", req.Id).Scan(&product.Id, &product.Name, &product.Description, &product.Price)
    if err != nil {
        log.Printf("Failed to get product: %v", err)
        return nil, err
    }
    return &product, nil
}

func (s *ProductServer) UpdateProduct(ctx context.Context, req *pb.UpdateProductRequest) (*pb.UpdateProductResponse, error) {
    _, err := s.db.Exec("UPDATE products SET name = $1, description = $2, price = $3 WHERE id = $4", req.Name, req.Description, req.Price, req.Id)
    if err != nil {
        log.Printf("Failed to update product: %v", err)
        return nil, err
    }
    return &pb.UpdateProductResponse{Id: req.Id}, nil
}

func (s *ProductServer) DeleteProduct(ctx context.Context, req *pb.DeleteProductRequest) (*pb.DeleteProductResponse, error) {
    _, err := s.db.Exec("DELETE FROM products WHERE id = $1", req.Id)
    if err != nil {
        log.Printf("Failed to delete product: %v", err)
        return nil, err
    }
    return &pb.DeleteProductResponse{Id: req.Id}, nil
}
