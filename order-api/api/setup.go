package api

import "github.com/gin-gonic/gin"

func Setup(engine *gin.Engine) {
	createOrder(engine)
}
