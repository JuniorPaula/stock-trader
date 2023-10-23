package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var port string

func main() {
	var err error

	if err = godotenv.Load(); err != nil {
		fmt.Println("Error loading .env file")
	}

	port = os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	fmt.Printf("Server is running on port %s\n", port)

	// start server
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: routes(),
	}
	err = srv.ListenAndServe()
	if err != nil {
		fmt.Print("ERROR: start server:\n", err)
		os.Exit(1)
	}
}
