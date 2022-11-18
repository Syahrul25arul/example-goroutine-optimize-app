package user_service

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/helper"
)

type UserService interface {
	Register(userRequest *domain.UserRegisterRequest) *helper.ResponseMessage
	VerifyEmail(UserVerifyEmailRequest *domain.UserVerifyEmailRequest) *helper.ResponseMessage
}
