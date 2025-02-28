package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/odd1n3rd/default_web/internal/config"
	"github.com/odd1n3rd/default_web/internal/entity"
	"github.com/odd1n3rd/default_web/internal/handler"
	"github.com/odd1n3rd/default_web/internal/repository"
	"github.com/odd1n3rd/default_web/internal/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	r := chi.NewRouter()

	cfg := config.Load()

	db, err := gorm.Open(postgres.Open(cfg.DBAddr), &gorm.Config{})
	if err != nil {
		log.Fatal(map[string]string{"error": err.Error()})
	}

	if err := db.AutoMigrate(&entity.User{}, &entity.Order{}); err != nil {
		log.Fatal(map[string]string{"error": err.Error()})
	}
	//user layer
	uRepo := repository.NewUserRepos(db)
	uService := service.NewUserService(uRepo)
	uHandler := handler.NewUserHandler(uService)

	//order layer
	oRepo := repository.NewOrderRepos(db)
	oService := service.NewOrderService(oRepo)
	oHandler := handler.NewOrderHandler(oService)

	r.Route("/users", func(r chi.Router) {
		r.Get("/", uHandler.GetUsers)
		r.Get("/{id}", uHandler.GetUser)
		r.Post("/", uHandler.CreateUser)
		r.Get("/{id}/orders", oHandler.GetOrders)
	})
	//uService := service.UserService(uRepo)
	// uHandler := handler.
	// uRepo.Create(&entity.User{Name: "hermiona", Email: "herm@hogwarts.com", Orders: []entity.Order{entity.Order{
	// 	ProductName: "philosoph stone",
	// 	Amount:      1,
	// 	Cost:        16,
	// }, entity.Order{
	// 	ProductName: "бузинная палочка",
	// 	Amount:      2,
	// 	Cost:        154,
	// }}})
	// uRepo.UpdateByID(&entity.User{Email: "whisley@hogwarts.com", Name: "ron"}, 1)
	// uRepo.DeleteByID(1)

	http.ListenAndServe(cfg.ServerAddress, r)
}
