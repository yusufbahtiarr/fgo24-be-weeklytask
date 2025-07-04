package controllers

import (
	"fgo24-be-ewallet/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
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

func GetAllHistory(ctx *gin.Context) {
	page, _ := strconv.Atoi(ctx.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "5"))
	userIdx, _ := ctx.Get("userId")
	userId := int(userIdx.(float64))

	history, err := models.FindHistoryTransaction(userId, page, limit)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Internal server error",
			Errors:  err.Error(),
		})
		return
	}

	totalData, err := models.GetTotalTransactionCount(userId)
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
		Message: "List History",
		PageInfo: map[string]int{
			"page":      page,
			"limit":     limit,
			"totalData": totalData,
		},
		Results: history,
	})
}

func UpdateProfile(ctx *gin.Context) {
	userIdx, _ := ctx.Get("userId")
	userId := int(userIdx.(float64))
	newData := models.UpdateProfileRq{}

	// println("ctr newData: ", newData)
	err := ctx.ShouldBind(&newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	err = models.UpdateProfile(newData, userId)
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

func UpdatePassword(ctx *gin.Context) {
	userIdx, _ := ctx.Get("userId")
	userId := int(userIdx.(float64))
	newData := models.Password{}

	err := ctx.ShouldBind(&newData)
	// fmt.Println("id: ", userId)
	// fmt.Println("ctrl: ", newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	err = models.UpdatePassword(newData, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed Update Password",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success Update Password",
	})
}

func UpdatePin(ctx *gin.Context) {
	userIdx, _ := ctx.Get("userId")
	userId := int(userIdx.(float64))
	newData := models.Pin{}
	err := ctx.ShouldBind(&newData)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid input",
		})
		return
	}

	err = models.UpdatePin(newData, userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed Update Pin",
			Errors:  err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Success Update Pin",
	})
}

func GetUserByName(ctx *gin.Context) {
	searchQy := ctx.Query("search")

	users, err := models.FindUserByName(searchQy)
	if err != nil {
		if err == pgx.ErrNoRows {
			ctx.JSON(http.StatusNotFound, models.Response{
				Success: false,
				Message: "no users matching the search criteria",
			})
			return
		}

		ctx.JSON(http.StatusInternalServerError, models.Response{
			Success: false,
			Message: "Failed to search users by name",
			Errors:  err.Error(),
		})
		return
	}

	var userViews []models.UserView
	for _, u := range users {
		view := models.UserView{}

		if u.Fullname != nil {
			view.FullName = *u.Fullname
		}
		if u.Phone != nil {
			view.Phone = *u.Phone
		}
		if u.ProfileImage != nil {
			view.ProfileImage = *u.ProfileImage
		}

		userViews = append(userViews, view)
	}

	ctx.JSON(http.StatusOK, models.Response{
		Success: true,
		Message: "Users found by name",
		Results: userViews,
	})
}

func GetBalance(ctx *gin.Context) {
	userIdx, _ := ctx.Get("userId")

	userId := int(userIdx.(float64))

	users, err := models.GetBalance(userId)
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
		Message: "Balance User",
		Results: users,
	})
}
