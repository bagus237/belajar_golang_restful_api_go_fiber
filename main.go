package main

import (
	"bagus/belajar_golang_restful_api/handler"
	"bagus/belajar_golang_restful_api/infrastructure"
	"bagus/belajar_golang_restful_api/internal/domain/category/repository"
	"bagus/belajar_golang_restful_api/internal/domain/category/service"
	"bagus/belajar_golang_restful_api/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	// Connect to the database
	db, err := infrastructure.NewDBConnection()
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Initialize repositories
	categoryRepository := repository.NewRepository(db)

	// Initialize services
	categoryService := service.NewCategoryService(categoryRepository)

	// Initialize handlers
	categoryHandler := handler.NewCategoryHandler(categoryService)

	// Initialize your router
	router := router.NewRouter(&router.DomainHandler{
		CategoryHandlerInterface: categoryHandler,
	})

	// Set up routes
	router.Route(app)

	// Start the server
	err = app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
