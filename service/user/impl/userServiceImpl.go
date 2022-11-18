package user_service_impl

import (
	"goroutine-optimize/domain"
	"goroutine-optimize/helper"
	repository_token "goroutine-optimize/repository/token"
	user_repo "goroutine-optimize/repository/user"
	"goroutine-optimize/utility"
	"net/http"
	"time"
)

type userServiceImpl struct {
	repo  user_repo.RepositoryUser
	token repository_token.RepositoryToken
}

func NewUserService(repo user_repo.RepositoryUser, repoToken repository_token.RepositoryToken) *userServiceImpl {
	return &userServiceImpl{repo: repo, token: repoToken}
}

func (s *userServiceImpl) Register(userRequest *domain.UserRegisterRequest) *helper.ResponseMessage {
	// check data user is valid
	if err := helper.IsValid(*userRequest); err != nil {
		// if data is not valid, return error message
		return helper.NewResponseMessage(true, err.Code, map[string]string{
			"message": err.Message,
		})
	} else {
		// bycript password
		userRequest.Password = helper.BcryptPassword(userRequest.Password)
		user := userRequest.ToUser()
		user.EmailVerify = "inactive"

		// if not any error, save data user
		if err := s.repo.Save(&user); err != nil {
			return helper.NewResponseMessage(true, err.Code, map[string]string{
				"message": err.Message,
			})
		} else {
			// create token for email verifiaction
			var token *domain.Token
			DateCreated := time.Now().Unix()
			tokenString := helper.RandomString(64)
			token = &domain.Token{
				Email:       user.Email,
				Token:       tokenString,
				DateCreated: DateCreated,
			}

			// save data token and check there is error or not
			if err = s.token.Save(token); err != nil {
				return helper.NewResponseMessage(true, err.Code, map[string]string{
					"message": err.Message,
				})
			}

			// end email to verify email_address
			go utility.SendEmailV2(user.Email, tokenString)
			// send response success
			return helper.NewResponseMessage(true, http.StatusOK, map[string]string{
				"message": "success registration",
			})
		}
	}
}

func (s *userServiceImpl) VerifyEmail(UserVerifyEmailRequest *domain.UserVerifyEmailRequest) *helper.ResponseMessage {
	// get data token by email
	if token, err := s.token.FindByEmail(UserVerifyEmailRequest.Email); err != nil {
		return helper.NewResponseMessage(true, err.Code, map[string]string{
			"message": err.Message,
		})
	} else {
		// cek token same with token in database
		if token.Token != UserVerifyEmailRequest.Token {
			return helper.NewResponseMessage(true, http.StatusForbidden, map[string]string{
				"message": "you dont have credential",
			})
		}

		// cek token has expired or not
		if time.Now().Unix() > (token.DateCreated + (60 * 60 * 24)) {
			return helper.NewResponseMessage(true, http.StatusForbidden, map[string]string{
				"message": "your token has expired",
			})
		}

		if user, err := s.repo.FindByEmail(token.Email); err != nil {
			return helper.NewResponseMessage(true, err.Code, map[string]string{
				"message": err.Message,
			})
		} else {
			// update user email verify
			user.EmailVerify = "active"
			if err := s.repo.Save(user); err != nil {
				return helper.NewResponseMessage(true, err.Code, map[string]string{
					"message": err.Message,
				})
			} else {
				return helper.NewResponseMessage(false, http.StatusOK, map[string]string{
					"message": "Your Email has been verify",
				})
			}
		}

	}
}
