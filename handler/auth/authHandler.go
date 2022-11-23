package auth_handler

import (
	"fmt"
	"goroutine-optimize/domain"
	"goroutine-optimize/helper"
	"goroutine-optimize/logger"
	service_user "goroutine-optimize/service/user"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type authHandler struct {
	service service_user.UserService
}

func NewAuthHandler(service service_user.UserService) *authHandler {
	return &authHandler{service: service}
}

func (h *authHandler) Register(ctx *gin.Context) {
	// * start time to record duration of execute process coding
	start := time.Now()

	// * create user variabel for data user
	var user *domain.UserRegisterRequest
	err := ctx.ShouldBindJSON(&user)

	// TODO : cek if any error when scan data user from client
	if err != nil {
		logger.Error("error scan data user")
		errResponse := helper.NewResponseMessage(true, http.StatusUnprocessableEntity, map[string]string{
			"message": "error scan data user" + err.Error(),
			"start":   fmt.Sprintf("Duration execution %s", time.Since(start)),
		})
		ctx.JSON(errResponse.Code, errResponse)
		return
	}

	// TODO : if not error, save data user and return response
	resp := h.service.Register(user)
	message := resp.Data.(map[string]string)
	message["start"] = fmt.Sprintf("Duration execution %s", time.Since(start))
	resp.Data = message

	ctx.JSON(resp.Code, resp)

}

func (a *authHandler) VerifyEmail(ctx *gin.Context) {
	// * start time to record duration of execute process coding
	start := time.Now()

	// TODO : create user variabel for data user email verify request
	var user = &domain.UserVerifyEmailRequest{}
	user.Email = ctx.Query("email")
	user.Token = ctx.Query("token")

	// TODO : Verify Email address and return response
	resp := a.service.VerifyEmail(user)
	message := resp.Data.(map[string]string)
	message["start"] = fmt.Sprintf("Duration execution %s", time.Since(start))
	resp.Data = message

	ctx.JSON(resp.Code, resp)

}
