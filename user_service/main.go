package user_service

import (
    "log"
    "net"

    "github.com/RImangali/user_service/handlers"
    "github.com/RImangali/user_service/pb"
    "google.golang.org/grpc"
)

func main() {
    lis, err := net.Listen("tcp", ":50051")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterUserServiceServer(s, &handlers.UserServer{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
