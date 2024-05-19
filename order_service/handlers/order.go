package handlers

import (
    "context"
    "database/sql"
    "log"
    "github.com/RImangali/order_service/pb"
    _ "github.com/lib/pq"
)

type OrderServer struct {
    pb.UnimplementedOrderServiceServer
}

var db *sql.DB

func init() {
    var err error
    db, err = sql.Open("localhost", "user=postgres password=Imangali2004 dbname=postgres sslmode=disable")
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    query := "INSERT INTO orders (user_id, product_id, quantity, status) VALUES ($1, $2, $3, $4) RETURNING id"
    var id int
    err := db.QueryRow(query, req.UserId, req.ProductId, req.Quantity, req.Status).Scan(&id)
    if err != nil {
        return nil, err
    }
    return &pb.CreateOrderResponse{Id: string(id)}, nil
}

func (s *OrderServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
    query := "SELECT id, user_id, product_id, quantity, status FROM orders WHERE id = $1"
    var id int
    var userId, productId, status string
    var quantity int
    err := db.QueryRow(query, req.Id).Scan(&id, &userId, &productId, &quantity, &status)
    if err != nil {
        return nil, err
    }
    return &pb.GetOrderResponse{Id: string(id), UserId: userId, ProductId: productId, Quantity: quantity, Status: status}, nil
}

func (s *OrderServer) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
    query := "UPDATE orders SET status = $1 WHERE id = $2 RETURNING id"
    var id int
    err := db.QueryRow(query, req.Status, req.Id).Scan(&id)
    if err != nil {
        return nil, err
    }
    return &pb.UpdateOrderResponse{Id: string(id)}, nil
}

func (s *OrderServer) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
    query := "DELETE FROM orders WHERE id = $1 RETURNING id"
    var id int
    err := db.QueryRow(query, req.Id).Scan(&id)
    if err != nil {
        return nil, err
    }
    return &pb.DeleteOrderResponse{Id: string(id)}, nil
}
