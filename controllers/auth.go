package controllers

import (
	"fgo24-be-ewallet/models"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/joho/godotenv"
)

func AuthLogin(ctx *gin.Context) {
	godotenv.Load()
	secretKey := os.Getenv("APP_SECRET")
	loginUser := models.LoginUser{}

	if err := ctx.ShouldBind(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	fmt.Println("controller:", loginUser.Email)
	user, err := models.FindUserByEmail(loginUser.Email)
	if err != nil {
		fmt.Println(user)
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "User with specified email not found",
		})
		return
	}
	fmt.Println("loginUser :", loginUser.Password)
	fmt.Println("User :", user.Password)
	if user == (models.User{}) || loginUser.Password != user.Password {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"iat":    time.Now().Unix(),
	})

	token, _ := generatedToken.SignedString([]byte(secretKey))

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success Login.",
		Results: map[string]string{
			"token": token,
		},
	})
}

func AuthRegister(ctx *gin.Context) {
	user := models.User{}

	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Input",
		})
		return
	}

	fmt.Println("ctr auth:", user)
	err = models.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Failed to create user.",
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Success created user.",
		Results: user.Email,
	})
}
