package main

import (
	"fgo24-be-ewallet/routers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	if os.Getenv("MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	// r.Use(middlewares.Cors())

	routers.CombineRouter(r)

	godotenv.Load()

	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
