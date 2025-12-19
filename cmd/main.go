package main

import (
	"E-Commerce/internal/config"
	"E-Commerce/internal/db"
	"E-Commerce/internal/handler"
	"E-Commerce/internal/model"
	"E-Commerce/internal/repository"
	"E-Commerce/internal/router"
	"E-Commerce/internal/service"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	cfg := config.Load()
	fmt.Println(cfg)
	dbConn, err := db.Connect()
	if err != nil {
		log.Fatalf("database connection failed: %v", err)
	}
	log.Println("database connected:", dbConn != nil)

	err = dbConn.AutoMigrate(
		&model.User{},
		&model.Product{},
		&model.Category{},
		&model.Cart{},
		&model.CartItem{},
		&model.Order{},
		&model.OrderItem{},
	)
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New(fiber.Config{
		AppName: "Arshia E-Commerce",
	})

	userRepo := repository.NewUserRepository(dbConn)
	authService := service.NewAuthService(userRepo, cfg.JWTSecret)
	authHandler := handler.NewAuthHandler(authService, userRepo)

	router.Setup(app, router.Deps{
		AuthHandler: authHandler,
		JWTSecret:   cfg.JWTSecret,
	})

	log.Fatal(app.Listen(":" + cfg.AppPort))
}
