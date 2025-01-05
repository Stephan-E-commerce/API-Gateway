package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/stepundel1/E-commerce/Users/logic/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	flag.Parse()

	var (
		nameEnter     string
		emailEnter    string
		passwordEnter string
	)
	// entering data for account
	fmt.Println("Name:")
	fmt.Scanln(&nameEnter)

	fmt.Println("Email:")
	fmt.Scanln(&emailEnter)

	fmt.Println("Password:")
	fmt.Scanln(&passwordEnter)

	var (
		name     = flag.String("name", nameEnter, "Name to greet")
		email    = flag.String("email", emailEnter, "Email of the user")
		password = flag.String("password", passwordEnter, "Password of the user")
	)

	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	// Вызовите метод регистрации пользователя
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	registerUser, err := c.RegisterUser(ctx, &pb.RegisterUserRequest{
		Name:     *name,
		Email:    *email,
		Password: *password,
	})
	if err != nil {
		log.Fatalf("could not register user: %v", err)
	}

	log.Printf("User registration success: %v", registerUser.GetSuccess())
}
