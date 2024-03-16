package application

import (
	"fmt"
	"net/http"

	"github.com/martinezhenry/search-products-dummy/internal/handlers"
)

// Run execute the application
func Run() {

	http.HandleFunc("/ping", handlers.PingHandler)
	http.HandleFunc("/pong", handlers.PongHandler)
	fmt.Println("running app...")
	http.ListenAndServe(":8080", nil)

}
