package user_service

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/helper"
)

type UserService interface {
	Save(user *domain.User) *helper.ResponseMessage
}
