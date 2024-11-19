package main

import (
	"fmt"
	"net/http"

	"github.com/oduortoni/art-pieces/controllers"
	utils "github.com/oduortoni/art-pieces/lib"
)

func main() {
	port := utils.Port()
	fmt.Printf("Server listening on http://localhost:%d\n", port)
	portStr := fmt.Sprintf("0.0.0.0:%d", port)

	http.HandleFunc("/", controllers.Index)
	http.ListenAndServe(portStr, nil)
}
