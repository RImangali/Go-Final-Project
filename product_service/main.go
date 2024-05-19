package product_service

import (
    "log"
    "net"

    "github.com/RImangali/product_service/handlers"
    "github.com/RImangali/product_service/pb"
    "google.golang.org/grpc"
)

func main() {
    lis, err := net.Listen("tcp", ":50052")
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }
    s := grpc.NewServer()
    pb.RegisterProductServiceServer(s, &handlers.ProductServer{})
    if err := s.Serve(lis); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
