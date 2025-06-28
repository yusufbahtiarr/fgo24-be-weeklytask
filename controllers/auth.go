package controllers

import (
	"fgo24-be-ewallet/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthLogin(ctx *gin.Context) {
	//inisiasi
	// godotenv.Load()
	// secretKey := os.Getenv("APP_SECRET")
	loginUser := models.LoginUser{}

	//cek data request
	if err := ctx.ShouldBind(&loginUser); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	//cek user by email
	user, err := models.FindUserByEmail(loginUser.Email)
	if err != nil {
		ctx.JSON(http.StatusNotFound, models.Response{
			Success: false,
			Message: "User with specified email not found",
		})
		return
	}

	//cek validasi password
	if user == (models.User{}) || loginUser.Password == user.Password {
		ctx.JSON(http.StatusUnauthorized, models.Response{
			Success: false,
			Message: "Invalid email or password",
		})
		return
	}

	//generate token
	// generatedToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
	// 	"userId": user.ID,
	// 	"iat":    time.Now().Unix(),
	// })

	// token, _ := generatedToken.SignedString([]byte(secretKey))

	//response success
	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success Login.",
		// Results: map[string]string{
		// 	"token": token,
		// },
	})
}

func AuthRegister(ctx *gin.Context) {
	//inisiasi
	user := models.User{}

	//cek data
	err := ctx.ShouldBind(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Input",
		})
		return
	}

	//proses register user
	fmt.Println("ctr auth:", user)
	err = models.RegisterUser(user)
	if err != nil {
		ctx.JSON(http.StatusOK, models.Response{
			Success: false,
			Message: "Failed to create user.",
		})
		return
	}

	//response success
	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Success created user.",
		Results: user.Email,
	})
}
