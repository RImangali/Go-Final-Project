package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/RImangali/product_service/pb"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("Did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewProductServiceClient(conn)

    switch os.Args[1] {
    case "create":
        createProduct(client)
    case "get":
        getProduct(client)
    case "update":
        updateProduct(client)
    case "delete":
        deleteProduct(client)
    default:
        fmt.Println("Unknown command")
    }
}

func createProduct(client pb.ProductServiceClient) {
    req := &pb.CreateProductRequest{Name: "Apple", Description: "Aport Apple from Almaty", Price: 120.0}
    res, err := client.CreateProduct(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not create product: %v", err)
    }
    fmt.Printf("Product created with ID: %s\n", res.Id)
}

func getProduct(client pb.ProductServiceClient) {
    req := &pb.GetProductRequest{Id: "1"}
    res, err := client.GetProduct(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not get product: %v", err)
    }
    fmt.Printf("Product: %s, Name: %s, Description: %s, Price: %.2f\n", res.Id, res.Name, res.Description, res.Price)
}

func updateProduct(client pb.ProductServiceClient) {
    req := &pb.UpdateProductRequest{Id: "1", Name: "Updated Product", Description: "This is an updated product", Price: 150.0}
    res, err := client.UpdateProduct(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not update product: %v", err)
    }
    fmt.Printf("Product updated with ID: %s\n", res.Id)
}

func deleteProduct(client pb.ProductServiceClient) {
    req := &pb.DeleteProductRequest{Id: "1"}
    res, err := client.DeleteProduct(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not delete product: %v", err)
    }
    fmt.Printf("Product deleted with ID: %s\n", res.Id)
}
