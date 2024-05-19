package handlers

import (
    "context"
    "database/sql"
    "log"
    "github.com/RImangali/user_service/pb"
    _ "github.com/lib/pq"
)

type UserServer struct {
    pb.UnimplementedUserServiceServer
}

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("localhost", "user=postgres password=Imangali2004 dbname=postgres sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    query := "INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id"
    var id int
    err := db.QueryRow(query, req.Name, req.Email, req.Password).Scan(&id)
    if err != nil {
        return nil, err
    }
    return &pb.CreateUserResponse{Id: string(id)}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    query := "SELECT id, name, email FROM users WHERE id = $1"
    var id int
    var name, email string
    err := db.QueryRow(query, req.Id).Scan(&id, &name, &email)
    if err != nil {
        return nil, err
    }
    return &pb.GetUserResponse{Id: string(id), Name: name, Email: email}, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
    query := "UPDATE users SET name = $1, email = $2 WHERE id = $3 RETURNING id"
    var id int
    err := db.QueryRow(query, req.Name, req.Email, req.Id).Scan(&id)
    if err != nil {
        return nil, err
    }
    return &pb.UpdateUserResponse{Id: string(id)}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
    query := "DELETE FROM users WHERE id = $1 RETURNING id"
    var id int
    err := db.QueryRow(query, req.Id).Scan(&id)
    if err != nil {
        return nil, err
    }
    return &pb.DeleteUserResponse{Id: string(id)}, nil
}
