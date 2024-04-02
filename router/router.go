package router

import (
	"bagus/belajar_golang_restful_api/handler"

	"github.com/gofiber/fiber/v2"
)

type DomainHandler struct {
	handler.CategoryHandlerInterface
}

type RouterStruct struct {
	DomainHandler DomainHandler
}

func NewRouter(handler *DomainHandler) RouterStruct {
	return RouterStruct{
		DomainHandler: *handler,
	}
}

func (router *RouterStruct) Route(app *fiber.App) {

	app.Get("/api/categories", router.DomainHandler.CategoryHandlerInterface.FindAllCategory)
	app.Get("/api/categories/:categoryId", router.DomainHandler.CategoryHandlerInterface.FindByIdCategory)
	app.Post("/api/categories", router.DomainHandler.CategoryHandlerInterface.CreateCategory)
	app.Put("/api/categories/:categoryId", router.DomainHandler.CategoryHandlerInterface.UpdateCategory)
	app.Delete("/api/categories/:categoryId", router.DomainHandler.CategoryHandlerInterface.DeleteCategory)
}
