# Go-Final-Project

## First, set-up your database in main.go files of each microservice(order, product, user)

## Second turn on the server
enter in a terminal 'go run main.go' of each microservice(order, product, user)

## Proceed to a cmd folder of any microservice type
enter in a terminal 'cd/microservice-type-you-desire/cmd'

## CRUD Operations
go run user_cli.go create
go run user_cli.go get 1
go run user_cli.go update 1
go run user_cli.go delete 1

go run order_cli.go create
go run order_cli.go get 1
go run order_cli.go update 1
go run order_cli.go delete 1

go run product_cli.go create
go run product_cli.go get 1
go run product_cli.go update 1
go run product_cli.go delete 1
