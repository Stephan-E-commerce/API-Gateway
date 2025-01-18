package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/stepundel1/E-commerce/API-Gateway/controller"
	pb "github.com/stepundel1/E-commerce/API-Gateway/proto/users"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewGreeterClient(conn)
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		controller.RegisterUser(w, r, c)
	})
	http.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		controller.LoginUser(w, r, c)
	})

	fmt.Println("Run in http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
