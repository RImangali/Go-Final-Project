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

	"github.com/RImangali/order_service/pb" // Update with your actual import path
)

type server struct {
	db *gorm.DB
	pb.UnimplementedOrderServiceServer
}

// Order represents the order model
type Order struct {
	ID        string `gorm:"primary_key"`
	UserID    string
	ProductID string
	Quantity  int32
}

func (s *server) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	order := &Order{
		UserID:    req.UserId,
		ProductID: req.ProductId,
		Quantity:  req.Quantity,
	}
	if err := s.db.Create(order).Error; err != nil {
		return nil, err
	}
	return &pb.CreateOrderResponse{Id: order.ID}, nil
}

func (s *server) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	var order Order
	if err := s.db.First(&order, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}
	return &pb.GetOrderResponse{Id: order.ID, UserId: order.UserID, ProductId: order.ProductID, Quantity: order.Quantity}, nil
}

// HTTP Handlers
func (s *server) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	var req pb.CreateOrderRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := s.CreateOrder(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (s *server) handleGetOrder(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	req := pb.GetOrderRequest{Id: id}
	resp, err := s.GetOrder(context.Background(), &req)
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

	db.AutoMigrate(&Order{})

	grpcLis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, &server{db: db})
	reflection.Register(grpcServer)

	go func() {
		fmt.Println("gRPC server is running on port :50053")
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	httpServer := &server{db: db}

	http.HandleFunc("/order", httpServer.handleCreateOrder)
	http.HandleFunc("/getorder", httpServer.handleGetOrder)

	fmt.Println("HTTP server is running on port :8082")
	if err := http.ListenAndServe(":8082", nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
