package main

import ("fmt"
		"net/http"
		"practice/internal/handlers"
		"practice/internal/config"
		"practice/internal/db"
		"context"
)

func main(){
	http.HandleFunc("/", handlers.HelloGo)
	fmt.Println("Sta99900")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("GetSuck")
	}
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil{
		fmt.Println("gg")
	}
	pool, err:= db.NewPool(ctx, cfg.Database.URL)
	if err != nil {
		fmt.Println("gg")
	}
	defer pool.Close()
	
}