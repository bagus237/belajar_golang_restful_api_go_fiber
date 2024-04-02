package infrastructure

import (
	"bagus/belajar_golang_restful_api/internal/domain/category/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// NewDBConnection membuat koneksi baru ke database
func NewDBConnection() (*gorm.DB, error) {
	dsn := "host=localhost user=postgres password=admin dbname=belajar_golang_restful_api port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Add Auto Migrate
	err = db.AutoMigrate(
		model.Category{})

	if err != nil {
		return nil, err
	}

	return db, nil

}
