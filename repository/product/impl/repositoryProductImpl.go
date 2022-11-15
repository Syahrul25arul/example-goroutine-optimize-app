package repository_product_impl

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/errs"
	"goroutine-optimize/logger"

	"gorm.io/gorm"
)

type repositoryProductImpl struct {
	db *gorm.DB
}

func NewRepositoryProduct(db *gorm.DB) *repositoryProductImpl {
	return &repositoryProductImpl{db: db}
}

func (r *repositoryProductImpl) SaveProduct(product *domain.Product) *errs.AppErr {
	if tx := r.db.Save(product); tx.Error != nil {
		logger.Error("error save product : " + tx.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return nil
}
