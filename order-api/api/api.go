package api

import (
	"net/http"

	"order-now/order-api/service"

	"github.com/gin-gonic/gin"
)

func createOrder(engine *gin.Engine) {

	engine.POST("createorder/", func(c *gin.Context) {

		response := service.ProcessOrder(c.Request.Body)

		c.JSON(http.StatusOK, response)
	})

}
