package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"

	"grpc_client/shopping"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(5*time.Second))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := shopping.NewShoppingServiceClient(conn)

	product := &shopping.ProductRequest{
		Name:     "Apples",
		Quantity: 10,
	}

	response, err := client.AddProduct(context.Background(), product)
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	fmt.Println(response.Message)

	productList, err := client.ListProducts(context.Background(), &shopping.Void{})
	if err != nil {
		log.Fatalf("Could not list products: %v", err)
	}

	fmt.Println("Product List:")
	for _, p := range productList.Products {
		fmt.Printf("Name: %s, Quantity: %d, Purchased: %v\n", p.Name, p.Quantity, p.Purchased)
	}
}
