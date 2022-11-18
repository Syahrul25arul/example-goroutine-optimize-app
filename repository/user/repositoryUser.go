package repository_user

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/errs"
)

type RepositoryUser interface {
	Save(user *domain.User) *errs.AppErr
	FindByEmail(email string) (*domain.User, *errs.AppErr)
}
