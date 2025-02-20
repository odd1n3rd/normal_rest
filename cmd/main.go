package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/odd1n3rd/default_web/internal/config"
	"github.com/odd1n3rd/default_web/internal/entity"
	"github.com/odd1n3rd/default_web/internal/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := chi.NewRouter()

	cfg := config.Load()

	db, err := gorm.Open(postgres.Open("host=localhost user=postgres dbname=clear_arch sslmode=disable password=' '"), &gorm.Config{})
	if err != nil {
		log.Fatal(map[string]string{"error": err.Error()})
	}

	if err := db.AutoMigrate(&entity.User{}, &entity.Order{}); err != nil {
		log.Fatal(map[string]string{"error": err.Error()})
	}

	uRepo := repository.NewUserRepos(db)
	// uService := service.UserService(uRepo)
	// uHandler := handler.
	//uRepo.Create(&entity.User{Name: "harry", Email: "harry@hogwarts.com"})

	fmt.Print(uRepo.GetByID(1))

	http.ListenAndServe(cfg.ServerAddress, r)
}
