package handler

import (
	"bagus/belajar_golang_restful_api/internal/domain/category/model/dto"
	"bagus/belajar_golang_restful_api/internal/domain/category/service"

	"github.com/gofiber/fiber/v2"
)

type CategoryHandlerImpl struct {
	CategoryService service.CategoryService
}

type CategoryHandlerInterface interface {
	CreateCategory(c *fiber.Ctx) (err error)
	UpdateCategory(c *fiber.Ctx) (err error)
	DeleteCategory(c *fiber.Ctx) (err error)
	FindByIdCategory(c *fiber.Ctx) (err error)
	FindAllCategory(c *fiber.Ctx) (err error)
}

func NewCategoryHandler(categoryService service.CategoryService) *CategoryHandlerImpl {
	return &CategoryHandlerImpl{
		CategoryService: categoryService}

}

func (handler *CategoryHandlerImpl) CreateCategory(c *fiber.Ctx) (err error) {

	var request dto.CategoryCreateRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	//category := dto.CategoryCreateRequestToCategory(&request)

	category := handler.CategoryService.Save(c.Context(), &request)
	return c.JSON(category)

}

func (handler *CategoryHandlerImpl) UpdateCategory(c *fiber.Ctx) (err error) {
	var request dto.CategoryUpdateRequest

	if err := c.BodyParser(&request); err != nil {
		return err
	}

	category := handler.CategoryService.Update(c.Context(), &request)
	return c.JSON(category)
}

func (handler *CategoryHandlerImpl) DeleteCategory(c *fiber.Ctx) (err error) {

	var request dto.CategoryUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	handler.CategoryService.Delete(c.Context(), request.Id)
	return nil

}

func (handler *CategoryHandlerImpl) FindByIdCategory(c *fiber.Ctx) (err error) {

	var request dto.CategoryUpdateRequest
	if err := c.BodyParser(&request); err != nil {
		return err
	}

	category := handler.CategoryService.FindById(c.Context(), request.Id)
	return c.JSON(category)
}

func (handler *CategoryHandlerImpl) FindAllCategory(c *fiber.Ctx) (err error) {

	categories := handler.CategoryService.FindAll(c.Context())
	return c.JSON(categories)
}
