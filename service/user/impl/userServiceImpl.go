package user_service_impl

import (
	"context"
	"goroutine-optimize/domain"
	"goroutine-optimize/errs"
	"goroutine-optimize/helper"
	"goroutine-optimize/logger"
	repository_token "goroutine-optimize/repository/token"
	user_repo "goroutine-optimize/repository/user"
	"goroutine-optimize/utility"
	"net/http"
	"sync"
	"time"

	"github.com/thanhpk/randstr"
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
	// if err := helper.IsValid(*userRequest); err != nil {
	// 	// if data is not valid, return error message
	// 	return helper.NewResponseMessage(true, err.Code, map[string]string{
	// 		"message": err.Message,
	// 	})
	// } else {
	// 	// TODO : Cek email has been used
	// 	if _, err := s.repo.FindUserAccount(userRequest.Username, userRequest.Email); err == nil {
	// 		return helper.NewResponseMessage(true, http.StatusBadRequest, map[string]string{
	// 			"message": "Username or Email has been registered",
	// 		})
	// 	}

	// 	// bycript password
	// 	userRequest.Password = helper.BcryptPassword(userRequest.Password)
	// 	user := userRequest.ToUser()
	// 	user.EmailVerify = "inactive"

	// 	// if not any error, save data user
	// 	if err := s.repo.Save(&user); err != nil {
	// 		return helper.NewResponseMessage(true, err.Code, map[string]string{
	// 			"message": err.Message,
	// 		})
	// 	} else {
	// 		// create token for email verifiaction
	// 		var token *domain.Token
	// 		DateCreated := time.Now().Unix()
	// 		tokenString := randstr.Hex(16)
	// 		token = &domain.Token{
	// 			Email:       user.Email,
	// 			Token:       tokenString,
	// 			DateCreated: DateCreated,
	// 		}

	// 		// save data token and check there is error or not
	// 		if err = s.token.Save(token); err != nil {
	// 			return helper.NewResponseMessage(true, err.Code, map[string]string{
	// 				"message": err.Message,
	// 			})
	// 		}

	// 		// end email to verify email_address
	// 		go utility.SendEmailV2(user.Email, tokenString)
	// 		// send response success
	// 		return helper.NewResponseMessage(false, http.StatusOK, map[string]string{
	// 			"message": "success registration",
	// 		})
	// 	}
	// }

	// * create wait group for waiting goroutine for this method.
	var wg sync.WaitGroup
	// * create channel
	errCh := make(chan *errs.AppErr)

	// * create context with cancel for cancel goroutine
	ctxParent := context.Background()
	ctx, cancel := context.WithCancel(ctxParent)

	// * add number running gorotuine for validation request client
	wg.Add(1)
	go helper.IsValidV2(*userRequest, errCh, ctx)

	// TODO : cek if any goroutine send error then cancel all goroutine
	if err := <-errCh; err != nil {
		if err.Message != "End Of Line" {
			cancel()
			return helper.NewResponseMessage(true, err.Code, map[string]string{
				"message": err.Message,
			})
		} else {
			wg.Done()
			close(errCh)
		}
	}
	cancel()
	wg.Wait()
	// TODO : Cek email has been used
	if _, err := s.repo.FindUserAccount(userRequest.Username, userRequest.Email); err == nil {
		return helper.NewResponseMessage(true, http.StatusBadRequest, map[string]string{
			"message": "Username or Email has been registered",
		})
	}

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
		tokenString := randstr.Hex(16)
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
		return helper.NewResponseMessage(false, http.StatusOK, map[string]string{
			"message": "success registration",
		})
	}
}

func (s *userServiceImpl) VerifyEmail(UserVerifyEmailRequest *domain.UserVerifyEmailRequest) *helper.ResponseMessage {
	// TODO: get data token by email
	if token, err := s.token.FindByEmail(UserVerifyEmailRequest.Email); err != nil {
		return helper.NewResponseMessage(true, err.Code, map[string]string{
			"message": err.Message,
		})
	} else {
		// TODO : cek token same with token in database
		if token.Token != UserVerifyEmailRequest.Token {
			logger.Warning("token not same")
			return helper.NewResponseMessage(true, http.StatusForbidden, map[string]string{
				"message": "you dont have credential",
			})
		}

		// cek token has expired or not
		if time.Now().Unix() > (token.DateCreated + (60 * 60 * 24)) {
			// TODO : delete token expired
			if err := s.token.Delete(token); err != nil {
				return helper.NewResponseMessage(true, http.StatusInternalServerError, map[string]string{
					"message": "something went wrong, please try again or request new email verification",
				})
			}
			return helper.NewResponseMessage(true, http.StatusForbidden, map[string]string{
				"message": "your token has expired",
			})
		}

		// cek email has resgitered or not
		if user, err := s.repo.FindByEmail(token.Email); err != nil {
			// TODO : delete token has been used
			if err := s.token.Delete(token); err != nil {
				return helper.NewResponseMessage(true, http.StatusInternalServerError, map[string]string{
					"message": "something went wrong, please try again or request new email verification",
				})
			}
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
				// TODO : delete token has been used
				if err := s.token.Delete(token); err != nil {
					return helper.NewResponseMessage(true, http.StatusInternalServerError, map[string]string{
						"message": "something went wrong, please try again or request new email verification",
					})
				}
				return helper.NewResponseMessage(false, http.StatusOK, map[string]string{
					"message": "Your Email has been verify",
				})
			}
		}
	}
}
