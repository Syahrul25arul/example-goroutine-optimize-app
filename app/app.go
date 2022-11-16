package app

import (
	"fmt"
	"goroutine-optimize/config"
	"goroutine-optimize/database"
	handler_auth "goroutine-optimize/handler/auth"
	handler_product "goroutine-optimize/handler/product"
	repository_product "goroutine-optimize/repository/product/impl"
	repository_user "goroutine-optimize/repository/user/impl"
	service_product "goroutine-optimize/service/product/impl"
	service_user "goroutine-optimize/service/user/impl"

	"github.com/gin-gonic/gin"
)

func Start() {

	// prepare database
	db := database.GetClientDb()

	// prepare repository
	repositoryProduct := repository_product.NewRepositoryProduct(db)
	repositoryUser := repository_user.NewRepositoryUser(db)

	// prepare service
	serviceProduct := service_product.NewServiceProduct(repositoryProduct)
	serviceUser := service_user.NewUserService(repositoryUser)

	// prepare handler
	handlerProduct := handler_product.NewHandlerProduct(serviceProduct)
	handlerAuth := handler_auth.NewAuthHandler(serviceUser)

	// setup gin
	r := gin.Default()

	r.POST("/products", handlerProduct.SaveProduct)
	r.POST("/register", handlerAuth.Register)

	r.POST("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "hello world")
	})

	// run server
	r.Run(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT))
}
