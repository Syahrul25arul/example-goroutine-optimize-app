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
		logger.Error("error save user : " + tx.Error.Error())
		return errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return nil
}

func (r *repositoryUserImpl) FindByEmail(email string) (*domain.User, *errs.AppErr) {
	// set variable to get data user
	var user *domain.User
	// get user by email
	if tx := r.db.First(&user, "email = ?", email); tx.RowsAffected == 0 {
		logger.Error("error get data user by email not found")
		return nil, errs.NewForbiddenError("you dont have credential")
	} else if tx.Error != nil {
		logger.Error("error get data user by email : " + tx.Error.Error())
		return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return user, nil
}

func (r *repositoryUserImpl) FindUserAccount(username, email string) (*domain.User, *errs.AppErr) {
	// * set variable to get data user
	var user *domain.User
	// TODO : get user by email
	if tx := r.db.First(&user, "username = ? or email = ?", username, email); tx.RowsAffected == 0 {
		logger.Error("error get data user account not found")
		return nil, errs.NewNotFoundError("user not found")
	} else if tx.Error != nil {
		logger.Error("error get data user account : " + tx.Error.Error())
		return nil, errs.NewUnexpectedError("Sorry, an error has occurred on our system due to an internal server error. please try again!")
	}
	return user, nil
}
