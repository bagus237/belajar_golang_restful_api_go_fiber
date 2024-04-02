package dto

import "bagus/belajar_golang_restful_api/internal/domain/category/model"

type CategoryUpdateRequest struct {
	Id   int    `validate:"required"`
	Name string `validate:"required,max=200,min=1" json:"name"`
}

func CategoryUpdateRequestToCategory(request *CategoryUpdateRequest) model.Category {

	return model.Category{
		ID:   request.Id,
		Name: request.Name,
	}
}
