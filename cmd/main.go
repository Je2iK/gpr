package main

import (
	"fmt"
	"net/http"
	"practice/internal/handlers"
	"practice/internal/repository"
	"practice/internal/service"
	"context"
	"practice/internal/config"
	"practice/internal/db"
	"log"
	"practice/internal/migrations"
)

func main() {
	ctx := context.Background()
	cfg, err := config.Load()
	if err != nil{
		log.Fatalf("НЕ ЗАГРУЗИЛСЯ КОНФИГ: %v", err)
		return 
	}

	pool, err := db.NewPool(ctx, cfg.Database.URL)
	if err != nil{
		log.Fatal("Net connecta :w", err)
	}
	migrations.RunMigrations(cfg.Database.URL)
	
	userRepository:= repository.NewUserRepository(pool)
	userService:= service.NewUserService(userRepository)
	userHandler:=handlers.NewUserHandler(userService)
	http.HandleFunc("/users/add", userHandler.CreateUser)
	http.HandleFunc("/users/delete", userHandler.DeleteUser)
	http.HandleFunc("/", handlers.HelloGo)
	fmt.Println("Sta99900")

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("GetSuck")
	}
	
}
