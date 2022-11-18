package repository_token_impl

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/errs"
	"goroutine-optimize/logger"

	"gorm.io/gorm"
)

type repositoryTokenImpl struct {
	db *gorm.DB
}

func NewRepositoryToken(db *gorm.DB) *repositoryTokenImpl {
	return &repositoryTokenImpl{db: db}
}

func (r *repositoryTokenImpl) Save(token *domain.Token) *errs.AppErr {
	// save token and check there is an error or not
	if tx := r.db.Save(token); tx.Error != nil {
		// if there error, create logger to record error
		logger.Error("error save token : " + tx.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return nil
}

func (r *repositoryTokenImpl) FindByEmail(email string) (*domain.Token, *errs.AppErr) {
	// set variable to get data token
	var token *domain.Token
	// get user by email
	if tx := r.db.First(&token, "email = ?", email); tx.RowsAffected == 0 {
		logger.Error("error get data token by email not found")
		return nil, errs.NewForbiddenError("you dont have credential")
	} else if tx.Error != nil {
		logger.Error("error get data token by email : " + tx.Error.Error())
		return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return token, nil
}
