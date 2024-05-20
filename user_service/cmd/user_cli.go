package main

import (
    "context"
    "fmt"
    "log"
    "os"

    "github.com/RImangali/user_service/pb"
    "google.golang.org/grpc"
)

func main() {
    conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
    if err != nil {
        log.Fatalf("did not connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewUserServiceClient(conn)

    switch os.Args[1] {
    case "create":
        createUser(client)
    case "get":
        getUser(client)
    case "update":
        updateUser(client)
    case "delete":
        deleteUser(client)
    default:
        fmt.Println("Unknown command")
    }
}

func createUser(client pb.UserServiceClient) {
    req := &pb.CreateUserRequest{Name: "Azamat Kozhakov", Email: "210107116@stu.sdu.edu.kz", Password: "Azeke04"}
    res, err := client.CreateUser(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not create user: %v", err)
    }
    fmt.Printf("User created with ID: %s\n", res.Id)
}

func getUser(client pb.UserServiceClient) {
    req := &pb.GetUserRequest{Id: "1"}
    res, err := client.GetUser(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not get user: %v", err)
    }
    fmt.Printf("User: %s, Email: %s\n", res.Name, res.Email)
}

func updateUser(client pb.UserServiceClient) {
    req := &pb.UpdateUserRequest{Id: "2", Name: "Rakhmetolla Assar", Email: "210107159@stu.sdu.edu.kz"}
    res, err := client.UpdateUser(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not update user: %v", err)
    }
    fmt.Printf("User updated with ID: %s\n", res.Id)
}

func deleteUser(client pb.UserServiceClient) {
    req := &pb.DeleteUserRequest{Id: "7"}
    res, err := client.DeleteUser(context.Background(), req)
    if err != nil {
        log.Fatalf("Could not delete user: %v", err)
    }
    fmt.Printf("User deleted with ID: %s\n", res.Id)
}
