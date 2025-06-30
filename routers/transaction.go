package routers

import (
	"fgo24-be-ewallet/controllers"
	"fgo24-be-ewallet/middlewares"

	"github.com/gin-gonic/gin"
)

func transactionRouter(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.POST("/transfer", controllers.CreateTransactionTransfer)
	r.POST("/topup", controllers.CreateTransactionTopup)
}
