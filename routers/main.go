package routers

import "github.com/gin-gonic/gin"

func CombineRouter(r *gin.Engine) {
	userRouter(r.Group("/users"))
	authRouter(r.Group("/auth"))
}

// eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3NTA5MzUzOTIsInVzZXJJZCI6N30.IOV4e9LZDZoZ4Q-DMF2itBewSH3L48c46rd5UwJscY8
