package repository

import (
	"bagus/belajar_golang_restful_api/internal/domain/category/model"
	"context"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Save(ctx context.Context, category_model *model.Category) model.Category
	Update(ctx context.Context, category_model *model.Category) model.Category
	Delete(ctx context.Context, categoryId int) (*model.Category, error)
	FindById(ctx context.Context, categoryId int) model.Category
	FindAll(ctx context.Context) []model.Category
}

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepositoryImpl{
		db: db,
	}
}

func (repo *categoryRepositoryImpl) Save(ctx context.Context, category_model *model.Category) model.Category {
	err := repo.db.Create(category_model).Error
	if err != nil {
		panic(err)
	}
	return *category_model

}

func (repo *categoryRepositoryImpl) Update(ctx context.Context, category_model *model.Category) model.Category {
	err := repo.db.Save(category_model).Error
	if err != nil {
		panic(err)
	}
	return *category_model
}

func (repo *categoryRepositoryImpl) Delete(ctx context.Context, categoryId int) (*model.Category, error) {
	var category model.Category
	err := repo.db.Delete(&category, categoryId).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

func (repo *categoryRepositoryImpl) FindById(ctx context.Context, categoryId int) model.Category {
	var category model.Category
	err := repo.db.First(&category, categoryId).Error
	if err != nil {
		panic(err)
	}
	return category
}

func (repo *categoryRepositoryImpl) FindAll(ctx context.Context) []model.Category {
	var categories []model.Category
	err := repo.db.Find(&categories).Error
	if err != nil {
		panic(err)
	}
	return categories
}
