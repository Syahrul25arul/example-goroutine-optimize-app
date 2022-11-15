package repository_product

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/errs"
)

type RepositoryProduct interface {
	SaveProduct(product *domain.Product) *errs.AppErr
}
