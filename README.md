# Shopping Service - gRPC Project

This project implements a shopping list service using gRPC and Redis. It allows users to manage a shopping list with products that have names, quantities, and purchased status. The service supports adding, updating, deleting, listing products, and marking them as purchased.

## Features

- Add a product to the shopping list.
- Get details of a specific product.
- Update the quantity of a product.
- Delete a product from the list.
- List all products in the shopping list.
- Mark a product as purchased.

## Project Structure

- **`shopping.proto`**: Defines the gRPC service and messages used for communication between the client and server.
- **Server**: Implements the gRPC service using Redis for storing and managing shopping list products.
- **Client**: Communicates with the server to perform CRUD operations on the shopping list.

## Prerequisites

- [Go](https://golang.org/) installed (version 1.16+).
- [Docker](https://www.docker.com/) installed and running.
- [Redis](https://redis.io/) running in a Docker container.
- [Protocol Buffers](https://developers.google.com/protocol-buffers) (protoc) installed.

## Setup Instructions

### 1. Run Redis in Docker

```bash
docker run --name redis -d -p 6379:6379 redis
```

### 2. Generate gRPC code from Proto file

To generate Go code from the proto file, navigate to the directory containing `shopping.proto` and run:

```bash
protoc --go_out=. --go-grpc_out=. shopping.proto
```

This will generate the required Go files for the gRPC service.

### 3. Run the Server

To start the server, run the following command:

```bash
go run server/main.go
```

The server will be running on port `9090` and will use Redis to store the product data.

### 4. Run the Client

To interact with the server, run the client program:

```bash
go run client/main.go
```

The client will add, update, list, and delete products from the shopping list, showing the output in the terminal.

## Proto Definitions

The `shopping.proto` file defines the following messages and service:

- **Messages**:
    - `Product`: Represents a product with a name, quantity, and purchased status.
    - `ProductRequest`: Used to send product data to the server (name, quantity).
    - `ProductNameRequest`: Used to request operations on a specific product by name.
    - `ProductResponse`: Holds a response message for operations like add, update, and delete.
    - `ProductList`: Contains a list of `Product` objects.
    - `void`: An empty message used for list operations.

- **Service**: `ShoppingService`
    - `AddProduct(ProductRequest)`: Adds a product.
    - `GetProduct(ProductNameRequest)`: Retrieves a specific product.
    - `UpdateProduct(ProductRequest)`: Updates product details.
    - `DeleteProduct(ProductNameRequest)`: Deletes a product.
    - `ListProducts(void)`: Lists all products.
    - `MarkAsPurchased(ProductNameRequest)`: Marks a product as purchased.

## Example gRPC Operations

1. **Add Product**: Add a new product (e.g., Apples, Bananas).
2. **List Products**: Retrieve the list of products.
3. **Update Product**: Update the quantity of a product (e.g., change Apples' quantity to 15).
4. **Delete Product**: Remove a product from the list (e.g., Bananas).
5. **Mark As Purchased**: Mark a product as purchased (e.g., Oranges).
