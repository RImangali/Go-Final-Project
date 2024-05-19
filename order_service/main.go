package order_service

import (
    "log"
    "net"

    "github.com/RImangali/order_service/handlers"
    "github.com/RImangali/order_service/pb"
    "google.golang.org/grpc"
)

func main() {
    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterOrderServiceServer(s, &handlers.OrderServer{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
