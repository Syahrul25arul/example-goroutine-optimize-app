package app

import (
	"fmt"
	"goroutine-optimize/config"
	"goroutine-optimize/database"
	handler_auth "goroutine-optimize/handler/auth"
	handler_product "goroutine-optimize/handler/product"
	repository_product "goroutine-optimize/repository/product/impl"
	repository_token "goroutine-optimize/repository/token/impl"
	repository_user "goroutine-optimize/repository/user/impl"
	service_product "goroutine-optimize/service/product/impl"
	service_user "goroutine-optimize/service/user/impl"

	"github.com/gin-gonic/gin"
)

func Start() {
	// loading env variabel
	config.SetupEnv(".env")

	// check all variables are loaded
	config.SanityCheck()
	// check app run for develope or production
	// if os.Getenv("TESTING") == "true" {
	// 	// panic("this is end point")
	// 	helper.TruncateAllTable(dbClient)
	// 	database.SetupDataDummyTest(dbClient)
	// }

	// prepare database
	db := database.GetClientDb()

	// prepare repository
	repositoryProduct := repository_product.NewRepositoryProduct(db)
	repositoryUser := repository_user.NewRepositoryUser(db)
	repositoryToken := repository_token.NewRepositoryToken(db)

	// prepare service
	serviceProduct := service_product.NewServiceProduct(repositoryProduct)
	serviceUser := service_user.NewUserService(repositoryUser, repositoryToken)

	// prepare handler
	handlerProduct := handler_product.NewHandlerProduct(serviceProduct)
	handlerAuth := handler_auth.NewAuthHandler(serviceUser)

	// setup gin
	r := gin.Default()

	r.POST("/products", handlerProduct.SaveProduct)
	r.POST("/register", handlerAuth.Register)
	r.GET("/verify", handlerAuth.VerifyEmail)

	r.GET("/test", func(ctx *gin.Context) {
		test := ctx.Query("test")
		ctx.JSON(200, test)
	})

	// run server
	r.Run(fmt.Sprintf("%s:%s", config.SERVER_ADDRESS, config.SERVER_PORT))
}
