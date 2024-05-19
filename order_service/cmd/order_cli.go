package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/RImangali/order_service/pb"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewOrderServiceClient(conn)

    switch os.Args[1] {
    case "create":
        createOrder(client)
    case "get":
        getOrder(client)
    case "update":
        updateOrder(client)
    case "delete":
        deleteOrder(client)
    default:
        fmt.Println("Unknown command")
    }
}

func createOrder(client pb.OrderServiceClient) {
    req := &pb.CreateOrderRequest{UserId: "1", ProductId: "1", Quantity: 2, Status: "Pending"}
    res, err := client.CreateOrder(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not create order: %v", err)
    }
    fmt.Printf("Order created with ID: %s\n", res.Id)
}

func getOrder(client pb.OrderServiceClient) {
    req := &pb.GetOrderRequest{Id: "1"}
    res, err := client.GetOrder(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not get order: %v", err)
    }
    fmt.Printf("Order: %s, User: %s, Product: %s, Quantity: %d, Status: %s\n", res.Id, res.UserId, res.ProductId, res.Quantity, res.Status)
}

func updateOrder(client pb.OrderServiceClient) {
    req := &pb.UpdateOrderRequest{Id: "1", Status: "Shipped"}
    res, err := client.UpdateOrder(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not update order: %v", err)
    }
    fmt.Printf("Order updated with ID: %s\n", res.Id)
}

func deleteOrder(client pb.OrderServiceClient) {
    req := &pb.DeleteOrderRequest{Id: "1"}
    res, err := client.DeleteOrder(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not delete order: %v", err)
    }
    fmt.Printf("Order deleted with ID: %s\n", res.Id)
}
