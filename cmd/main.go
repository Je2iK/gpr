package main

import (
	"fmt"
	"net/http"
	"practice/internal/handlers"
)

func main() {
	http.HandleFunc("/", handlers.HelloGo)
	fmt.Println("Sta99900")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("GetSuck")
	}

}
