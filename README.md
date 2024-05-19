# Go-Final-Project

## Set-up your database in main.go files of each microservice(order, product, user)

## Create 3 tables(orders, products, users) in your database
+ CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    email VARCHAR(100) UNIQUE,
    password VARCHAR(100),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100),
    description TEXT,
    price DECIMAL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id),
    product_id INT REFERENCES products(id),
    quantity INT,
    status VARCHAR(50),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


## Turn on the server
enter in a terminal 'go run main.go' of each microservice(order, product, user)

## Proceed to a cmd folder of any microservice type
enter in a terminal 'cd/microservice-type-you-desire/cmd'

## CRUD Operations
##### go run user_cli.go create
##### go run user_cli.go get 1
##### go run user_cli.go update 1
##### go run user_cli.go delete 1

##### go run order_cli.go create
##### go run order_cli.go get 1
##### go run order_cli.go update 1
##### go run order_cli.go delete 1

##### go run product_cli.go create
##### go run product_cli.go get 1
##### go run product_cli.go update 1
##### go run product_cli.go delete 1
