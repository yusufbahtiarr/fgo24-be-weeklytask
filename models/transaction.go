package models

import (
	"context"
	"fgo24-be-ewallet/utils"
	"fmt"
	"time"
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

	trx, err := conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction")
	}
	defer trx.Rollback(context.Background())

	var currentBalance float64
	err = trx.QueryRow(
		context.Background(),
		"SELECT balance FROM users WHERE id = $1 FOR UPDATE",
		userId,
	).Scan(&currentBalance)

	if err != nil {
		return fmt.Errorf("failed get balance user: (%w)", err)
	}

	if currentBalance < float64(transaction.Amount) {
		return fmt.Errorf("insufficient balance")
	}

	_, err = trx.Exec(
		context.Background(),
		"UPDATE users SET balance = balance - $1 WHERE id = $2",
		transaction.Amount, userId,
	)

	if err != nil {
		return fmt.Errorf("failed to sender balance: (%w)", err)
	}

	_, err = trx.Exec(
		context.Background(),
		"UPDATE users SET balance = balance + $1 WHERE id = $2",
		transaction.Amount, transaction.ReceiverId,
	)

	if err != nil {
		return fmt.Errorf("failed to update receiver balance: (%w)", err)
	}

	_, err = trx.Exec(
		context.Background(),
		`INSERT INTO transactions (transaction_type, amount, description, sender_id, receiver_id) VALUES ($1, $2, $3, $4, $5)`,
		"transfer",
		transaction.Amount,
		transaction.Description,
		userId,
		transaction.ReceiverId,
	)

	if err != nil {
		return fmt.Errorf("failed to insert transaction: %w", err)
	}

	if err := trx.Commit(context.Background()); err != nil {
		return fmt.Errorf("failed to commit transaction: (%w)", err)
	}

	return nil
}

func CreateTransactionTopup(transaction TransactionTopup, userId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	trx, err := conn.Begin(context.Background())
	if err != nil {
		return fmt.Errorf("failed to begin transaction")
	}
	defer trx.Rollback(context.Background())

	_, err = trx.Exec(
		context.Background(),
		"UPDATE users SET balance = balance + $1 WHERE id = $2",
		transaction.Amount, userId,
	)

	if err != nil {
		return fmt.Errorf("failed to update balance: (%w)", err)
	}

	_, err = conn.Exec(
		context.Background(),
		`INSERT INTO transactions (transaction_type, amount, receiver_id, payment_method_id) VALUES ($1, $2, $3, $4)`,
		"topup",
		transaction.Amount,
		userId,
		transaction.PaymentMethodId,
	)

	if err != nil {
		return fmt.Errorf("failed to insert transaction: (%w)", err)
	}

	if err := trx.Commit(context.Background()); err != nil {
		return fmt.Errorf("failed to commit transaction: (%w)", err)
	}

	return nil
}
