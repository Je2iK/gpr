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
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"practice/internal/middlewares"
)

func main() {
	r:= chi.NewRouter()
	runmiddlewares(r)
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request){

	})
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
	r.Post("/users/add", userHandler.CreateUser)
	r.Delete("/users/delete", userHandler.DeleteUser)
	fmt.Println("Sta99900")

	http.ListenAndServe(":8080", r)

}
