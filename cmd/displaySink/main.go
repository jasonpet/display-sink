package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", httpHandler)

	// Start the server
	fmt.Println("Listening for events on port 8080")
	http.ListenAndServe(":8080", nil)
}

func httpHandler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println(request.Body)
	writer.WriteHeader(http.StatusOK)
}
