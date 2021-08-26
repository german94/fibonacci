package main

import (
	"fmt"
	"net/http"
)

func main() {
	startServer()
}

func startServer() {
	http.HandleFunc("/", calcHandler)
	fmt.Println("Server started at port 8080")
	http.ListenAndServe(":8080", nil)
}
