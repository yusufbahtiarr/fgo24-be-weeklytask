package models

import (
	"context"
	"fgo24-be-ewallet/utils"
	"time"

	"github.com/jackc/pgx/v5/pgconn"
)

type Transaction struct {
	ID              int       `db:"id" form:"id" json:"id"`
	TransactionType string    `db:"transaction_type" form:"transaction_type" json:"transaction_type" binding:"required,transaction_type"`
	Amount          int       `db:"amount" form:"amount" json:"amount" binding:"required,amount"`
	Description     *string   `db:"description" form:"description" json:"description" binding:"required,description"`
	SenderId        *int      `db:"sender_id" form:"sender_id" json:"sender_id" binding:"required,sender_id"`
	ReceiverId      *int      `db:"receiver_id" form:"receiver_id" json:"receiver_id" binding:"required,receiver_id"`
	PaymentMethodId *int      `db:"payment_method_id" form:"payment_method_id" json:"payment_method_id" binding:"required,payment_method_id"`
	CreatedAt       time.Time `db:"created_at" form:"created_at" json:"created_at" binding:"required,created_at"`
}

type TransactionTransfer struct {
	Amount      int    `form:"amount" json:"amount" binding:"required"`
	Description string `form:"description"`
	ReceiverId  int    `form:"receiver_id" binding:"required"`
}
type TransactionTopup struct {
	Amount          int `form:"amount" json:"amount" binding:"required"`
	PaymentMethodId int `form:"payment_method_id" binding:"required"`
}

type Transactions []Transaction

func CreateTransactionTransfer(transaction TransactionTransfer, userId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()
	transfer := "transfer"

	_, err = conn.Exec(
		context.Background(),
		`INSERT INTO transactions (transaction_type, amount, description, sender_id, receiver_id) VALUES ($1, $2, $3, $4, $5)`,
		transfer,
		transaction.Amount,
		transaction.Description,
		userId,
		transaction.ReceiverId,
	)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return err
		}
		return err
	}

	return nil
}

func CreateTransactionTopup(transaction TransactionTopup, userId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()
	topup := "topup"

	_, err = conn.Exec(
		context.Background(),
		`INSERT INTO transactions (transaction_type, amount, receiver_id, payment_method_id) VALUES ($1, $2, $3, $4)`,
		topup,
		transaction.Amount,
		userId,
		transaction.PaymentMethodId,
	)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return err
		}
		return err
	}

	return nil
}
