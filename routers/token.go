package routers

import (
	"fgo24-be-ewallet/controllers"

	"github.com/gin-gonic/gin"
)

func tokenRouter(r *gin.RouterGroup) {
	r.GET("", controllers.GenerateToken)
	r.POST("", controllers.VerifyToken)
}
