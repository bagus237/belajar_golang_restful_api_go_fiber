package dto

import "bagus/belajar_golang_restful_api/internal/domain/category/model"

type CategoryResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func CategoryToCategoryResponse(category *model.Category) CategoryResponse {
	return CategoryResponse{
		Id:   category.ID,
		Name: category.Name,
	}
}
