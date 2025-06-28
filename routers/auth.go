package routers

import (
	"fgo24-be-ewallet/controllers"

	"github.com/gin-gonic/gin"
)

func authRouter(r *gin.RouterGroup) {
	r.POST("/login", controllers.AuthLogin)
	r.POST("/register", controllers.AuthRegister)
}
