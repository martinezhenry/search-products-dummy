package application

import (
	"net/http"

	"github.com/martinezhenry/search-products-dummy/internal/handlers"
)

// Run execute the application
func Run() {
	
	http.HandleFunc("/ping", handlers.PingHandler)
	http.ListenAndServe(":8080", nil)

}
