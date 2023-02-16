package main

import (
	"fmt"
	"net/http"

	"github.com/SophisticaSean/flight_path_calculator/internal/controllers"
)

func main() {
	http.Handle("/calculate", http.HandlerFunc(controllers.CalculateHandler))
	fmt.Println("listening on localhost:8080/calculate")
	// ignoring the error value returned by ListenAndServe
	_ = http.ListenAndServe(":8080", nil)
}
