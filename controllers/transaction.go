package controllers

import (
	"fgo24-be-ewallet/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateTransactionTransfer(ctx *gin.Context) {
	userIdx, _ := ctx.Get("userId")
	userId := int(userIdx.(float64))
	transaction := models.TransactionTransfer{}

	err := ctx.ShouldBind(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Input",
		})
		return
	}

	err = models.CreateTransactionTransfer(transaction, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Success created transaction transfer.",
	})
}

func CreateTransactionTopup(ctx *gin.Context) {
	userIdx, _ := ctx.Get("userId")
	userId := int(userIdx.(float64))
	transaction := models.TransactionTopup{}

	err := ctx.ShouldBind(&transaction)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: "Invalid Input",
		})
		return
	}

	err = models.CreateTransactionTopup(transaction, userId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, models.Response{
		Success: true,
		Message: "Success created transaction topup.",
	})
}
