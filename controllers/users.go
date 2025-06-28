package controllers

import (
	"fgo24-be-ewallet/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUser(ctx *gin.Context) {
	userIdx, _ := ctx.Get("userId")

	userId := int(userIdx.(float64))

	fmt.Printf("User yang sedang login adalah user dengan id %d\n", userId)

	users, err := models.FindAllUsers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "List User",
		Results: users,
	})
}

func GetUserByEmail(ctx *gin.Context) {
	user := models.User{}

	if err := ctx.ShouldBind(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Input",
		})
		return
	}
	fmt.Println("email: ", user.Email)
	user2, err := models.FindUserByEmail(user.Email)
	if err != nil {
		fmt.Println(user2)
		ctx.JSON(http.StatusOK, models.Response{
			Success: false,
			Message: "User email not found",
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "User by Email",
		Results: user2,
	})
}

func UpdateProfile(ctx *gin.Context) {
	newData := models.User{}
	err := ctx.ShouldBind(&newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	err = models.UpdateProfile(newData)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed Update User",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success Update User",
	})

}
