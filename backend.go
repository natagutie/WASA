package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gorilla/mux"
)

type User struct {
	Username   string `json:"username"`
	Identifier string `json:"identifier"`
}

var users = make(map[string]User)

func main() {
	r := mux.NewRouter()

	// Register API handlers here
	r.HandleFunc("/login", doLogin).Methods("POST")
	r.HandleFunc("/setMyUserName", setMyUserName).Methods("POST")
	r.HandleFunc("/uploadPhoto", uploadPhoto).Methods("POST")
	r.HandleFunc("/followUser", followUser).Methods("POST")
	r.HandleFunc("/unfollowUser", unfollowUser).Methods("POST")
	r.HandleFunc("/banUser", banUser).Methods("POST")
	r.HandleFunc("/unbanUser", unbanUser).Methods("POST")
	r.HandleFunc("/getUserProfile", getUserProfile).Methods("GET")
	r.HandleFunc("/getMyStream", getMyStream).Methods("GET")
	r.HandleFunc("/likePhoto", likePhoto).Methods("POST")
	r.HandleFunc("/unlikePhoto", unlikePhoto).Methods("POST")
	r.HandleFunc("/commentPhoto", commentPhoto).Methods("POST")
	r.HandleFunc("/uncommentPhoto", uncommentPhoto).Methods("POST")
	r.HandleFunc("/deletePhoto", deletePhoto).Methods("POST")

	// Start web server
	go func() {
		port := "8080" // Change to your desired port
		fmt.Printf("Server listening on :%s...\n", port)
		log.Fatal(http.ListenAndServe(":"+port, r))
	}()

	// Wait for termination signal
	waitForTermination()
}

func doLogin(w http.ResponseWriter, r *http.Request) {
	var userRequest map[string]string
	if err := json.NewDecoder(r.Body).Decode(&userRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	username, ok := userRequest["name"]
	if !ok || len(username) < 3 || len(username) > 16 {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	// Simplified logic: Generate a unique identifier for the user
	identifier := generateIdentifier()

	// Save user data
	users[identifier] = User{Username: username, Identifier: identifier}

	response := map[string]string{"identifier": identifier}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}

func setMyUserName(w http.ResponseWriter, r *http.Request) {
	var usernameRequest map[string]string
	if err := json.NewDecoder(r.Body).Decode(&usernameRequest); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	identifier, ok := getUsernameFromContext(r)
	if !ok {
		http.Error(w, "User not authenticated", http.StatusUnauthorized)
		return
	}

	newUsername, ok := usernameRequest["username"]
	if !ok {
		http.Error(w, "Invalid username", http.StatusBadRequest)
		return
	}

	// Update username
	users[identifier] = User{Username: newUsername, Identifier: identifier}

	w.WriteHeader(http.StatusOK)
}

func uploadPhoto(w http.ResponseWriter, r *http.Request) {
	// Implement your uploadPhoto logic here
	w.WriteHeader(http.StatusCreated)
}

// Implement other handler functions for each API endpoint

// Helper function to generate a unique identifier (simplified)
func generateIdentifier() string {
	return "abcdef012345"
}

// Helper function to get the user identifier from the request context (simplified)
func getUsernameFromContext(r *http.Request) (string, bool) {
	return "abcdef012345", true
}
