package handlers

import (
    "context"
    "database/sql"
    "log"

    "github.com/RImangali/order_service/pb"
)

type OrderServer struct {
    pb.UnimplementedOrderServiceServer
    db *sql.DB
}

func NewOrderServer(db *sql.DB) *OrderServer {
    return &OrderServer{db: db}
}

func (s *OrderServer) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
    var id string
    err := s.db.QueryRow("INSERT INTO orders (user_id, product_id, quantity, status) VALUES ($1, $2, $3, $4) RETURNING id", req.UserId, req.ProductId, req.Quantity, req.Status).Scan(&id)
    if err != nil {
        log.Printf("Failed to create order: %v", err)
        return nil, err
    }
    return &pb.CreateOrderResponse{Id: id}, nil
}

func (s *OrderServer) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
    var order pb.GetOrderResponse
    err := s.db.QueryRow("SELECT id, user_id, product_id, quantity, status FROM orders WHERE id = $1", req.Id).Scan(&order.Id, &order.UserId, &order.ProductId, &order.Quantity, &order.Status)
    if err != nil {
        log.Printf("Failed to get order: %v", err)
        return nil, err
    }
    return &order, nil
}

func (s *OrderServer) UpdateOrder(ctx context.Context, req *pb.UpdateOrderRequest) (*pb.UpdateOrderResponse, error) {
    _, err := s.db.Exec("UPDATE orders SET status = $1 WHERE id = $2", req.Status, req.Id)
    if err != nil {
        log.Printf("Failed to update order: %v", err)
        return nil, err
    }
    return &pb.UpdateOrderResponse{Id: req.Id}, nil
}

func (s *OrderServer) DeleteOrder(ctx context.Context, req *pb.DeleteOrderRequest) (*pb.DeleteOrderResponse, error) {
    _, err := s.db.Exec("DELETE FROM orders WHERE id = $1", req.Id)
    if err != nil {
        log.Printf("Failed to delete order: %v", err)
        return nil, err
    }
    return &pb.DeleteOrderResponse{Id: req.Id}, nil
}
