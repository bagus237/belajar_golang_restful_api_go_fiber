package service

import (
	"bagus/belajar_golang_restful_api/internal/domain/category/model"
	"bagus/belajar_golang_restful_api/internal/domain/category/model/dto"
	"bagus/belajar_golang_restful_api/internal/domain/category/repository"
	"context"
)

type CategoryService interface {
	Save(ctx context.Context, save_request *dto.CategoryCreateRequest) dto.CategoryResponse
	Update(ctx context.Context, category_model *dto.CategoryUpdateRequest) dto.CategoryResponse
	Delete(ctx context.Context, categoryId int) error
	FindById(ctx context.Context, categoryId int) error
	FindAll(ctx context.Context) []model.Category
}

type categoryServiceImpl struct {
	categoryRepository repository.CategoryRepository
}

func NewCategoryService(categoryRepository repository.CategoryRepository) CategoryService {
	return &categoryServiceImpl{
		categoryRepository: categoryRepository,
	}
}

func (service *categoryServiceImpl) Save(ctx context.Context, save_request *dto.CategoryCreateRequest) dto.CategoryResponse {

	category := service.categoryRepository.Save(ctx, &model.Category{
		Name: save_request.Name,
	})
	return dto.CategoryToCategoryResponse(&category)
}

func (service *categoryServiceImpl) Update(ctx context.Context, category_model *dto.CategoryUpdateRequest) dto.CategoryResponse {
	category := service.categoryRepository.Update(ctx, &model.Category{
		ID:   category_model.Id,
		Name: category_model.Name,
	})

	return dto.CategoryToCategoryResponse(&category)
}

func (service *categoryServiceImpl) Delete(ctx context.Context, categoryId int) error {

	service.categoryRepository.Delete(ctx, categoryId)
	return nil
}

func (service *categoryServiceImpl) FindById(ctx context.Context, categoryId int) error {
	service.categoryRepository.FindById(ctx, categoryId)
	return nil
}

func (service *categoryServiceImpl) FindAll(ctx context.Context) []model.Category {
	return service.categoryRepository.FindAll(ctx)
}
