package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/RImangali/product_service/pb" // Update with your actual import path
)

type server struct {
	db *gorm.DB
	pb.UnimplementedProductServiceServer
}

// Product represents the product model
type Product struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Price float64
}

func (s *server) CreateProduct(ctx context.Context, req *pb.CreateProductRequest) (*pb.CreateProductResponse, error) {
	product := &Product{
		Name:  req.Name,
		Price: req.Price,
	}
	if err := s.db.Create(product).Error; err != nil {
		return nil, err
	}
	return &pb.CreateProductResponse{Id: product.ID}, nil
}

func (s *server) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	var product Product
	if err := s.db.First(&product, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}
	return &pb.GetProductResponse{Id: product.ID, Name: product.Name, Price: product.Price}, nil
}

// HTTP Handlers
func (s *server) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	var req pb.CreateProductRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := s.CreateProduct(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (s *server) handleGetProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	req := pb.GetProductRequest{Id: id}
	resp, err := s.GetProduct(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func main() {
	dbHost := "localhost"
	dbUser := "postgres"
	dbPassword := "Imangali2004"
	dbName := "postgres"

	dsn := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, dbUser, dbName, dbPassword)
	db, err := gorm.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	db.AutoMigrate(&Product{})

	grpcLis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterProductServiceServer(grpcServer, &server{db: db})
	reflection.Register(grpcServer)

	go func() {
		fmt.Println("gRPC server is running on port :50052")
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	httpServer := &server{db: db}

	http.HandleFunc("/product", httpServer.handleCreateProduct)
	http.HandleFunc("/getproduct", httpServer.handleGetProduct)

	fmt.Println("HTTP server is running on port :8081")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
