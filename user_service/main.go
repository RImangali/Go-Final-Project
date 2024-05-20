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

	"github.com/RImangali/user_service/pb" // Update with your actual import path
)

type server struct {
	db *gorm.DB
	pb.UnimplementedUserServiceServer
}

// User represents the user model
type User struct {
	ID    string `gorm:"primary_key"`
	Name  string
	Email string
    Password string
}

func (s *server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user := &User{
		Name:  req.Name,
		Email: req.Email,
        Password: req.Password,
	}
	if err := s.db.Create(user).Error; err != nil {
		return nil, err
	}
	return &pb.CreateUserResponse{Id: user.ID}, nil
}

func (s *server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	var user User
	if err := s.db.First(&user, "id = ?", req.Id).Error; err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{Id: user.ID, Name: user.Name, Email: user.Email}, nil
}

// HTTP Handlers
func (s *server) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	var req pb.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	resp, err := s.CreateUser(context.Background(), &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(resp)
}

func (s *server) handleGetUser(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "id parameter is required", http.StatusBadRequest)
		return
	}

	req := pb.GetUserRequest{Id: id}
	resp, err := s.GetUser(context.Background(), &req)
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

	db.AutoMigrate(&User{})

	grpcLis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &server{db: db})
	reflection.Register(grpcServer)

	go func() {
		fmt.Println("gRPC server is running on port :50051")
		if err := grpcServer.Serve(grpcLis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	httpServer := &server{db: db}

	http.HandleFunc("/user", httpServer.handleCreateUser)
	http.HandleFunc("/getuser", httpServer.handleGetUser)

	fmt.Println("HTTP server is running on port :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}
}
