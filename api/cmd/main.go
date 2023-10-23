package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Server is running on port 8080")

	// start server
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}
	err := srv.ListenAndServe()
	if err != nil {
		fmt.Print("ERROR: start server:\n", err)
		os.Exit(1)
	}
}
