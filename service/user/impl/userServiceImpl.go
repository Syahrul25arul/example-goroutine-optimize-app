package user_service_impl

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/helper"
	user_repo "goroutine-optimize/repository/user"
	"goroutine-optimize/utility"
	"net/http"
)

type userServiceImpl struct {
	repo user_repo.RepositoryUser
}

func NewUserService(repo user_repo.RepositoryUser) *userServiceImpl {
	return &userServiceImpl{repo: repo}
}

func (s *userServiceImpl) Save(user *domain.User) *helper.ResponseMessage {
	// check data user is valid
	if err := helper.IsValid(*user); err != nil {
		// if data is not valid, return error message
		return helper.NewResponseMessage(true, err.Code, map[string]string{
			"message": err.Message,
		})
	} else {
		// bycript password
		user.Password = helper.BcryptPassword(user.Password)
		// if not any error, save data user
		if err := s.repo.Save(user); err != nil {
			return helper.NewResponseMessage(true, err.Code, map[string]string{
				"message": err.Message,
			})
		} else {
			// end email to verify email_address
			go utility.SendEmail(user.Email)
			// send response success
			return helper.NewResponseMessage(true, http.StatusOK, map[string]string{
				"message": "success registration",
			})
		}
	}
}
