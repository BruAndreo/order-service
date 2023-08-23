package db

import (
	"log"

	"github.com/bruandreo/order-service/internal/domain"
	"gorm.io/gorm"
)

type ProductRepositoryDatabase struct {
	DB *gorm.DB
}

func (prd *ProductRepositoryDatabase) Create(product domain.Product) bool {
	result := prd.DB.Unscoped().Create(&product)
	if result.Error != nil {
		log.Print("ERROR", result.Error.Error())
		return false
	}

	return true
}
