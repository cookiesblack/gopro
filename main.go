package main

import (
	"fmt"
	"net/http"

	"gopro/main/routes"

	"github.com/gorilla/mux"
)

func main() {
	// Create a new instance of the mux router
	router := mux.NewRouter()

	// Register routes from another file
	routes.RegisterRoutes(router)

	// Start the server
	fmt.Println("Server started on port 8080")
	http.ListenAndServe(":8080", router)
}
