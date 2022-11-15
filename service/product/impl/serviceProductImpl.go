package service_product_impl

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/helper"
	repository_product "goroutine-optimize/repository/product"
)

type serviceProductImpl struct {
	repo repository_product.RepositoryProduct
}

func NewServiceProduct(repo repository_product.RepositoryProduct) *serviceProductImpl {
	return &serviceProductImpl{repo: repo}
}

func (service *serviceProductImpl) Save(product *domain.Product) *helper.ResponseMessage {
	// if err := helper.IsValid(product); err != nil {
	// 	return helper.NewResponseMessage(true, err.Code, map[string]string{"message": err.Message})
	// }
	if resp := service.repo.SaveProduct(product); resp != nil {
		return helper.NewResponseMessage(true, resp.Code, map[string]string{"message": resp.Message})
	} else {
		return helper.NewResponseMessage(false, 200, map[string]string{
			"message": "product has been created",
		})
	}
}
