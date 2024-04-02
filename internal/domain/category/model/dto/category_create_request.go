package dto

import "bagus/belajar_golang_restful_api/internal/domain/category/model"

type CategoryCreateRequest struct {
	Name string `validate:"required,min=1,max=100" json:"name"`
}

func CategoryCreateRequestToCategory(request *CategoryCreateRequest) model.Category {

	return model.Category{
		Name: request.Name,
	}
}
