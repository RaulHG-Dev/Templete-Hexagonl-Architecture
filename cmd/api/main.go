package main

import (
	"architecture-hexagonal/handlers/user"
	"architecture-hexagonal/internal/repositories/mysql"
	user3 "architecture-hexagonal/internal/repositories/mysql/user"
	user2 "architecture-hexagonal/internal/services/user"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	// This is the main entry point for the API service.
	// The implementation will depend on the specific requirements of your application.
	// For example, you might set up an HTTP server, define routes, etc.

	// Here, we are loading environment variables from a .env file.
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}

	// Initialize the database connection and other services here.
	r := mux.NewRouter()
	var conn mysql.DB

	repository := user3.Repository{
		Client: conn.GetInstance(),
	}

	userService := &user2.Service{
		Repo: repository,
	}

	userHandler := &user.Handler{
		UserService: userService,
	}

	r.HandleFunc("/users", userHandler.CreateUser).Methods("GET")

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
