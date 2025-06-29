package routers

import (
	"fgo24-be-ewallet/controllers"
	"fgo24-be-ewallet/middlewares"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.GET("", controllers.GetAllUser)
	r.GET("/", controllers.SearchUserByName)
	r.GET("/history", controllers.GetAllHistory)
	r.POST("/email", controllers.GetUserByEmail)
	r.POST("/transfer", controllers.CreateTransactionTransfer)
	r.POST("/topup", controllers.CreateTransactionTopup)
	r.PATCH("/profile", controllers.UpdateProfile)
	r.PATCH("/password", controllers.UpdatePassword)
	r.PATCH("/pin", controllers.UpdatePin)
}
