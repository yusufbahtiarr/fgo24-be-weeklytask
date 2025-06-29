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
	r.POST("/email", controllers.GetUserByEmail)
	r.POST("/transfer", controllers.GetUserByEmail)
	r.POST("/topup", controllers.GetUserByEmail)
	r.PATCH("/profile", controllers.UpdateProfile)
}
