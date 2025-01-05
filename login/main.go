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
		emailEnterLogin    string
		passwordEnterLogin string
	)
	// entering data for account

	fmt.Println("Email:")
	fmt.Scanln(&emailEnterLogin)

	fmt.Println("Password:")
	fmt.Scanln(&passwordEnterLogin)

	var (
		email    = flag.String("email", emailEnterLogin, "Email of the user")
		password = flag.String("password", passwordEnterLogin, "Password of the user")
	)

	conn, err := grpc.NewClient("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	loginUser, err := c.LoginUser(ctx, &pb.LoginUserRequest{
		Email:    *email,
		Password: *password,
	})
	if err != nil {
		log.Fatalf("could not register user: %v", err)
	}

	log.Printf("User registration success: %v", loginUser.GetSuccess())
}
