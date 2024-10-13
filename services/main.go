package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/stepundel1/E-commerce/Users/logic/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	name = flag.String("name", "Steve", "Name to greet")
)
var (
	email = flag.String("email", "steve@mail.com", "Email of the user")
)
var (
	password = flag.String("password", "secretpassword", "Password of the user")
)

// func main() {
// 	flag.Parse()
// 	// Set up a connection to the server.
// 	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
// 	if err != nil {
// 		log.Fatalf("did not connect: %v", err)
// 	}
// 	defer conn.Close()
// 	c := pb.NewGreeterClient(conn)

// 	// Contact the server and print out its response.
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
// 	defer cancel()

// 	s, err := c.RegisterUser(ctx, &pb.RegisterUserRequest{Name: *name, Email: *email, Password: *password})
// 	if err != nil {
// 		log.Fatalf("could not greet: %v", err)
// 	}
// 	log.Printf("77: %s", s.GetSuccess())
// }

func main() {
	flag.Parse()

	// Установите соединение с gRPC сервером
	flag.Parse()
	// 	// Set up a connection to the server.
	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Вызовите метод регистрации пользователя
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.RegisterUser(ctx, &pb.RegisterUserRequest{
		Name:     *name,
		Email:    *email,
		Password: *password,
	})
	if err != nil {
		log.Fatalf("could not register user: %v", err)
	}

	log.Printf("User registration success: %v", resp.GetSuccess())
}
