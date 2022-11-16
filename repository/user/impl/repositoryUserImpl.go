package repository_user_impl

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/errs"
	"goroutine-optimize/logger"

	"gorm.io/gorm"
)

type repositoryUserImpl struct {
	db *gorm.DB
}

func NewRepositoryUser(db *gorm.DB) *repositoryUserImpl {
	return &repositoryUserImpl{db: db}
}

func (r *repositoryUserImpl) Save(user *domain.User) *errs.AppErr {
	// save user and check there is an error or not
	if tx := r.db.Save(user); tx.Error != nil {
		// if there error, create logger to record error
		logger.Error("error save product : " + tx.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return nil
}
