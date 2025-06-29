package models

import (
	"context"
	"fgo24-be-ewallet/utils"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Response struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Errors  any    `json:"errors,omitempty"`
	Results any    `json:"results,omitempty"`
}
type User struct {
	ID       int    `db:"id"  form:"id"`
	Email    string `db:"email" form:"email" json:"email" binding:"required,email"`
	Password string `db:"password" form:"password" json:"password"`
	Pin      string `db:"pin" form:"pin"`
	Fullname string `db:"fullname" form:"fullname"`
	Phone    string `db:"phone" form:"phone"`
}

type LoginUser struct {
	Email    string `db:"email" form:"email"`
	Password string `db:"password" form:"password"`
}

type Register struct {
	ID              int    `db:"id"  form:"id"`
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,password"`
	ConfirmPassword string `form:"confirm_password" binding:"required,confirm_password"`
}

type Password struct {
	Email           string `form:"email" binding:"required,email"`
	ExistPassword   string `form:"exist_password" binding:"required,exist_password"`
	NewPassword     string `form:"new_password" binding:"required,new_password"`
	ConfirmPassword string `form:"confirm_password" binding:"required,confirm_password"`
}

type Pin struct {
	ID    int    `db:"id"  form:"id"`
	Email string `form:"email" binding:"required,email"`
	Pin   string `form:"pin"`
}

type Transaction struct {
	ID              int    `db:"id"  form:"id"`
	TransactionType string `form:"transaction_type" binding:"required,transaction_type"`
	Amount          int    `form:"amount" binding:"required,amount"`
	Description     string `form:"description" binding:"required,description"`
	SenderId        string `form:"sender_id" binding:"required,sender_id"`
	ReceiverId      string `form:"receiver_id" binding:"required,receiver_id"`
	PaymentMethodId string `form:"payment_method_id" binding:"required,payment_method_id"`
}

type Users []User

func FindAllUsers() ([]User, error) {
	conn, err := utils.DBConnect()

	if err != nil {
		return []User{}, err
	}
	defer conn.Close()

	query := "SELECT id, email, password, pin, fullname, phone from users"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return []User{}, err
	}

	users, err := pgx.CollectRows[User](rows, pgx.RowToStructByName)
	if err != nil {
		return []User{}, err
	}
	return users, err
}

func FindUserByEmail(email string) (User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	fmt.Println("model:", email)
	query := `SELECT id, email, password, pin, fullname, phone FROM users WHERE email = $1`
	rows, err := conn.Query(context.Background(), query, email)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)
	if err != nil {
		return User{}, err
	}

	return user, err
}

func UpdateProfile(newData User) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	oldData, err := FindUserByEmail(newData.Email)
	if err != nil {
		return err
	}

	if newData.Email == "" && newData.Fullname == "" && newData.Phone == "" {
		return fmt.Errorf("input data must not be empty")
	}

	if newData.Email != "" && newData.Email != oldData.Email {
		oldData.Email = newData.Email
	}
	if newData.Fullname != "" && newData.Fullname != oldData.Fullname {
		oldData.Fullname = newData.Fullname
	}
	if newData.Phone != "" && newData.Phone != oldData.Phone {
		oldData.Phone = newData.Phone
	}

	_, err = conn.Exec(context.Background(), `UPDATE users set email =  $1, fullname = $2, phone = $3 where id=$4`, oldData.Email, oldData.Fullname, oldData.Phone, oldData.ID)

	return err
}

func UpdatePassword(newData Password) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	oldData, err := FindUserByEmail(newData.Email)
	if err != nil {
		return err
	}

	if newData.ExistPassword == "" && newData.NewPassword == "" && newData.ConfirmPassword == "" {
		return fmt.Errorf("input password cannot be empty")
	}

	if newData.NewPassword != newData.ConfirmPassword {
		return fmt.Errorf("new password and confirm password do not match")
	}

	if newData.NewPassword != oldData.Password {
		oldData.Password = newData.NewPassword
	}

	_, err = conn.Exec(context.Background(), `UPDATE users set password = $1 where id=$4`, oldData.Password, oldData.ID)

	return err
}

func UpdatePin(newData Pin) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	oldData, err := FindUserByEmail(newData.Email)
	if err != nil {
		return err
	}

	if newData.Pin == "" {
		return fmt.Errorf("input pin cannot be empty")
	}

	if newData.Pin != oldData.Pin {
		oldData.Pin = newData.Pin
	}

	_, err = conn.Exec(context.Background(), `UPDATE users set password =  $1 where id=$4`, oldData.Password, oldData.ID)

	return err
}

func SearchUserByName(name string) ([]Users, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	search := "%" + name + "%"
	query := `SELECT id, email, fullname, phone, balance, profile_iamge FROM users WHERE fullname ILIKE $1`
	rows, err := conn.Query(context.Background(), query, search)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	users, err := pgx.CollectRows[Users](rows, pgx.RowToStructByName)
	if err != nil {
		return nil, err
	}

	if len(users) == 0 {
		return []Users{}, nil
	}

	return users, nil
}

func FindHistoryTransaction(email string) (User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	query := `SELECT id, transaction_type, amount, description, created_at, sender_id, receiver_id, payment_method_id FROM transaction WHERE sender_id = $1`
	rows, err := conn.Query(context.Background(), query, email)
	if err != nil {
		return User{}, err
	}
	defer rows.Close()

	user, err := pgx.CollectOneRow[User](rows, pgx.RowToStructByName)
	if err != nil {
		return User{}, err
	}

	return user, err
}
