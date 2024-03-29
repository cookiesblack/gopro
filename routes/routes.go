package routes

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

// RegisterRoutes registers all routes
func RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/products", ProductsHandler).Methods("GET")
	router.HandleFunc("/articles", ArticlesHandler).Methods("GET")
	router.HandleFunc("/articles/{id}", ArticleHandler).Methods("GET")
}

// Handlers
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Home Page")
}

func ProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Products Page")
}

func ArticlesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Articles Page")
}

func ArticleHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintf(w, "Requested Article ID: %s", id)
}
