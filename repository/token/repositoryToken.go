package repository_token

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/errs"
)

type RepositoryToken interface {
	Save(token *domain.Token) *errs.AppErr
	FindByEmail(email string) (*domain.Token, *errs.AppErr)
	Delete(token *domain.Token) *errs.AppErr
}
