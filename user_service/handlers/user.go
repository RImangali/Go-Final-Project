package handlers

import (
    "context"
    "database/sql"
    "log"

    "github.com/RImangali/user_service/pb"
)

type UserServer struct {
    pb.UnimplementedUserServiceServer
    db *sql.DB
}

func NewUserServer(db *sql.DB) *UserServer {
    return &UserServer{db: db}
}

func (s *UserServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
    var id string
    err := s.db.QueryRow("INSERT INTO users (name, email, password) VALUES ($1, $2, $3) RETURNING id", req.Name, req.Email, req.Password).Scan(&id)
    if err != nil {
        log.Printf("Failed to create user: %v", err)
        return nil, err
    }
    return &pb.CreateUserResponse{Id: id}, nil
}

func (s *UserServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
    var user pb.GetUserResponse
    err := s.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", req.Id).Scan(&user.Id, &user.Name, &user.Email)
    if err != nil {
        log.Printf("Failed to get user: %v", err)
        return nil, err
    }
    return &user, nil
}

func (s *UserServer) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
    _, err := s.db.Exec("UPDATE users SET name = $1, email = $2 WHERE id = $3", req.Name, req.Email, req.Id)
    if err != nil {
        log.Printf("Failed to update user: %v", err)
        return nil, err
    }
    return &pb.UpdateUserResponse{Id: req.Id}, nil
}

func (s *UserServer) DeleteUser(ctx context.Context, req *pb.DeleteUserRequest) (*pb.DeleteUserResponse, error) {
    _, err := s.db.Exec("DELETE FROM users WHERE id = $1", req.Id)
    if err != nil {
        log.Printf("Failed to delete user: %v", err)
        return nil, err
    }
    return &pb.DeleteUserResponse{Id: req.Id}, nil
}
