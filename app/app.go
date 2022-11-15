package app

import (
	"fmt"
	"goroutine-optimize/database"
	handler_product "goroutine-optimize/handler/product"
	repository_product "goroutine-optimize/repository/product/impl"
	service_product "goroutine-optimize/service/product/impl"

	"github.com/gin-gonic/gin"
)

func Start() {

	// prepare database
	db := database.GetClientDb()

	// prepare repository
	repositoryProduct := repository_product.NewRepositoryProduct(db)

	// prepare service
	serviceProduct := service_product.NewServiceProduct(repositoryProduct)

	// prepare handler
	handlerProduct := handler_product.NewHandlerProduct(serviceProduct)

	// setup gin
	r := gin.Default()

	r.POST("/products", handlerProduct.SaveProduct)
	r.POST("/test", func(ctx *gin.Context) {
		ctx.JSON(200, "hello world")
	})

	// run server
	r.Run(fmt.Sprintf("%s:%s", "localhost", "9000"))
}
