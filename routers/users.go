package routers

import (
	"fgo24-be-ewallet/controllers"

	"github.com/gin-gonic/gin"
)

func userRouter(r *gin.RouterGroup) {
	r.GET("", controllers.GetAllUser)
	r.POST("/email", controllers.GetUserByEmail)
	r.POST("/transfer", controllers.GetUserByEmail)
	r.POST("/topup", controllers.GetUserByEmail)
	r.PATCH("/profile", controllers.UpdateProfile)
}
