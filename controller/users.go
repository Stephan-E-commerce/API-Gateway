package controller

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/stepundel1/E-commerce/API-Gateway/entity"
	pb "github.com/stepundel1/E-commerce/API-Gateway/proto/users"
)

func RegisterUser(w http.ResponseWriter, r *http.Request, client pb.GreeterClient) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// parsing data
	var req entity.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// sending data to gRPC server
	registerUser, err := client.RegisterUser(context.Background(), &pb.RegisterUserRequest{
		Name:     req.Username,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Printf("Error during user registration: %v", err)
		http.Error(w, "Failed to register user", http.StatusInternalServerError)
		return
	}

	// returning
	response := map[string]interface{}{
		"success": registerUser.GetSuccess(),
		"message": "Register user:success",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}

func LoginUser(w http.ResponseWriter, r *http.Request, client pb.GreeterClient) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// parsing data
	var req entity.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	// sending data to gRPC server
	loginUser, err := client.LoginUser(context.Background(), &pb.LoginUserRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		log.Printf("Error during user log in: %v", err)
		http.Error(w, "Failed to log in user", http.StatusInternalServerError)
		return
	}

	// returning
	response := map[string]interface{}{
		"success": loginUser.GetSuccess(),
		"message": "Log in:success",
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
