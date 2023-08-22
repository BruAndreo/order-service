package routes

import "github.com/bruandreo/order-service/internal/repositories"

type DependenciesInjector struct {
	ProductRepository repositories.ProductRepository
}
