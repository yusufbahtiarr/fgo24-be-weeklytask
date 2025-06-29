package main

import (
	"fgo24-be-ewallet/routers"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	r := gin.Default()

	if os.Getenv("MODE") == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	routers.CombineRouter(r)

	r.Run(fmt.Sprintf(":%s", os.Getenv("APP_PORT")))
}
