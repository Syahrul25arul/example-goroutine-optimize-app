package service_product

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/helper"
)

type ServiceProduct interface {
	Save(product *domain.Product) *helper.ResponseMessage
}
