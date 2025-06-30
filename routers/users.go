package routers

import (
	"fgo24-be-ewallet/controllers"
	"fgo24-be-ewallet/middlewares"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	r.Use(middlewares.VerifyToken())
	r.GET("", controllers.GetAllUser)
	r.GET("/search", controllers.GetUserByName)
	r.GET("/history", controllers.GetAllHistory)
	r.GET("/balance", controllers.GetBalance)
	r.POST("/email", controllers.GetUserByEmail)
	r.PATCH("/profile", controllers.UpdateProfile)
	r.PATCH("/password", controllers.UpdatePassword)
	r.PATCH("/pin", controllers.UpdatePin)
}
