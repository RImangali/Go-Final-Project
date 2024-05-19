package main

import (
    "database/sql"
    "log"
    "net"

    "github.com/RImangali/order_service/handlers"
    "github.com/RImangali/order_service/pb"
    "google.golang.org/grpc"
    _ "github.com/lib/pq" // Import the PostgreSQL driver
)

func main() {
    db, err := sql.Open("postgres", "user=postgres password=Imangali2004 dbname=postgres sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
    defer db.Close()

    // Ping the database to ensure connection
    if err := db.Ping(); err != nil {
        log.Fatalf("Failed to ping database: %v", err)
    }

    lis, err := net.Listen("tcp", ":50053")
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    grpcServer := grpc.NewServer()
    orderServer := handlers.NewOrderServer(db)
    pb.RegisterOrderServiceServer(grpcServer, orderServer)

    log.Println("Order service is running on port 50053")
    if err := grpcServer.Serve(lis); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
