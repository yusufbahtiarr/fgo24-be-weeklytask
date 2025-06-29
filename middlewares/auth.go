package middlewares

import (
	"fgo24-be-ewallet/models"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func VerifyToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		godotenv.Load()
		secretKey := os.Getenv("APP_SECRET")
		token := strings.Split(ctx.GetHeader("Authorization"), "Bearer")

		fmt.Println("token: ", token)
		if len(token) < 2 {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Success: false,
				Message: "Unauthorized",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}
		rawTokenString := strings.TrimSpace(token[1])
		rawToken, err := jwt.Parse(rawTokenString, func(t *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})

		fmt.Println("raw token: ", rawToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, models.Response{
				Success: false,
				Message: "Invalid token",
			})
			ctx.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		userId := rawToken.Claims.(jwt.MapClaims)["userId"]
		ctx.Set("userId", userId)
		ctx.Next()
	}
}
