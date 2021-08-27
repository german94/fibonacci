package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	startServer()
}

func startServer() {
	http.HandleFunc("/", calcHandler)
	fmt.Println("Server started at port 8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified
	}
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
