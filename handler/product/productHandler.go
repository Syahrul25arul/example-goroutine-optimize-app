package handler_product

import (
	"fmt"
	"goroutine-optimize/domain"
	"goroutine-optimize/helper"
	"goroutine-optimize/logger"
	serviceProduct "goroutine-optimize/service/product"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type handlerProduct struct {
	service serviceProduct.ServiceProduct
}

func NewHandlerProduct(service serviceProduct.ServiceProduct) *handlerProduct {
	return &handlerProduct{service: service}
}

func (h *handlerProduct) SaveProduct(ctx *gin.Context) {
	start := time.Now()

	// create variabel product and catch data product from request
	var product *domain.Product
	err := ctx.ShouldBindJSON(&product)

	if err != nil {
		logger.Error("error scan data product")
		errResponse := helper.NewResponseMessage(true, http.StatusUnprocessableEntity, map[string]string{"message": "error scan data product"})
		ctx.JSON(errResponse.Code, errResponse)
		return
	}

	resp := h.service.Save(product)
	message := resp.Data.(map[string]string)
	message["start"] = fmt.Sprintf("Duration executution %s", time.Since(start))
	resp.Data = message

	ctx.JSON(resp.Code, resp)

}
