package routes

import (
	"github.com/bruandreo/order-service/internal/repositories"
	"gorm.io/gorm"
)

type DependenciesInjector struct {
	ProductRepository repositories.ProductRepository
	DB                *gorm.DB
}
