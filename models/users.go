package models

import (
	"context"
	"fgo24-be-ewallet/utils"
	"fmt"

	"github.com/jackc/pgx/v5"
)

type Response struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Errors   any    `json:"errors,omitempty"`
	PageInfo any    `json:"pageinfo,omitempty"`
	Results  any    `json:"results,omitempty"`
}
type User struct {
	ID           int     `db:"id"  form:"id"`
	Email        string  `db:"email" form:"email" json:"email" binding:"required,email"`
	Password     string  `db:"password" form:"password" json:"password"`
	Pin          string  `db:"pin" form:"pin"`
	Fullname     *string `db:"fullname" form:"fullname"`
	Phone        *string `db:"phone" form:"phone"`
	ProfileImage *string `db:"profile_image" form:"profile_image"`
}

type UpdateProfileRq struct {
	Email        string `form:"email" json:"email" binding:"required,email"`
	Fullname     string `form:"fullname" json:"fullname"`
	Phone        string `form:"phone" json:"phone"`
	ProfileImage string `form:"profile_image" json:"profile_image"`
}

type LoginUser struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

type BalanceUser struct {
	ID       int     `db:"id" `
	Fullname *string `db:"fullname"`
	Balance  int     `db:"balance"`
}

type Register struct {
	Email           string `form:"email" binding:"required,email"`
	Password        string `form:"password" binding:"required,password"`
	ConfirmPassword string `form:"confirm_password" binding:"required,confirm_password"`
}

type Password struct {
	ExistPassword   string `form:"exist_password" binding:"required"`
	NewPassword     string `form:"new_password" binding:"required,min=8"`
	ConfirmPassword string `form:"confirm_password" binding:"required,eqfield=NewPassword"`
}

type Pin struct {
	Pin string `form:"pin" binding:"required,len=6"`
}

type Users []User

func FindAllUsers() (Users, error) {
	conn, err := utils.DBConnect()

	if err != nil {
		return Users{}, err
	}
	defer conn.Close()

	query := "SELECT id, email, password, pin, fullname, phone, profile_image from users"
	rows, err := conn.Query(context.Background(), query)
	if err != nil {
		return Users{}, err
	}

	users, err := pgx.CollectRows[User](rows, pgx.RowToStructByName)
	if err != nil {
		return Users{}, err
	}
	return users, err
}

func FindUserByEmail(email string) (User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	query := `SELECT id, email, password, pin, fullname, phone, profile_image FROM users WHERE email = $1`
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

func FindUserByID(id int) (User, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return User{}, err
	}
	defer conn.Close()

	// fmt.Println("model: ", id)
	query := `SELECT id, email, password, pin, fullname, phone, profile_image FROM users WHERE id = $1`
	rows, err := conn.Query(context.Background(), query, id)
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

func UpdateProfile(newData UpdateProfileRq, userId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	oldData, err := FindUserByID(userId)
	if err != nil {
		return err
	}

	if newData.Email == "" && newData.Fullname == "" && newData.Phone == "" {
		return fmt.Errorf("input data must not be empty")
	}

	if newData.Email != "" && newData.Email != oldData.Email {
		oldData.Email = newData.Email
	}
	if newData.Fullname != "" && newData.Fullname != *oldData.Fullname {
		oldData.Fullname = &newData.Fullname
	}
	if newData.Phone != "" && newData.Phone != *oldData.Phone {
		oldData.Phone = &newData.Phone
	}

	_, err = conn.Exec(context.Background(), `UPDATE users set email =  $1, fullname = $2, phone = $3 where id=$4`, oldData.Email, oldData.Fullname, oldData.Phone, oldData.ID)

	return err
}

func UpdatePassword(newData Password, userId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	oldData, err := FindUserByID(userId)
	if err != nil {
		return err
	}
	// fmt.Println("oldData:", oldData)
	// fmt.Println("newData:", newData)

	if newData.ExistPassword == "" && newData.NewPassword == "" && newData.ConfirmPassword == "" {
		return fmt.Errorf("input password cannot be empty")
	}

	if oldData.Password != newData.ExistPassword {
		return fmt.Errorf("existing password is incorrect")
	}

	if newData.NewPassword != newData.ConfirmPassword {
		return fmt.Errorf("new password and confirm password do not match")
	}

	_, err = conn.Exec(context.Background(), `UPDATE users set password = $1 where id=$2`, newData.NewPassword, userId)

	return err
}

func UpdatePin(newData Pin, userId int) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	oldData, err := FindUserByID(userId)
	if err != nil {
		return err
	}

	if newData.Pin == "" {
		return fmt.Errorf("input pin cannot be empty")
	}

	if newData.Pin != oldData.Pin {
		oldData.Pin = newData.Pin
	}

	_, err = conn.Exec(context.Background(), `UPDATE users set password =  $1 where id=$2`, oldData.Password, userId)

	return err
}

func FindUserByName(name string) ([]Users, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	search := "%" + name + "%"
	query := `SELECT id, email, fullname, phone, balance, profile_image FROM users WHERE fullname ILIKE $1`
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

func FindHistoryTransaction(userId, page, limit int) (Transactions, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return Transactions{}, err
	}
	defer conn.Close()

	offset := (page - 1) * limit
	if offset < 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 5
	}
	query := `SELECT id, transaction_type, amount, description, created_at, sender_id, receiver_id, payment_method_id FROM transactions WHERE sender_id = $3 or receiver_id = $3 ORDER BY created_at DESC LIMIT $1 OFFSET $2 `
	rows, err := conn.Query(context.Background(), query, limit, offset, userId)
	if err != nil {
		return Transactions{}, err
	}
	defer rows.Close()

	transaction, err := pgx.CollectRows[Transaction](rows, pgx.RowToStructByName)
	if err != nil {
		return Transactions{}, err
	}

	return transaction, err
}

func GetTotalTransactionCount(userId int) (int, error) {
	conn, err := utils.DBConnect()
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	var count int
	err = conn.QueryRow(context.Background(), "SELECT COUNT(*) FROM transactions WHERE sender_id = $1 or receiver_id = $1;", userId).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to get total transaction count: %w", err)
	}
	return count, nil
}

func GetBalance(userId int) (BalanceUser, error) {
	conn, err := utils.DBConnect()

	if err != nil {
		return BalanceUser{}, err
	}
	defer conn.Close()

	query := `SELECT id, fullname, balance FROM users WHERE id = $1`
	rows, err := conn.Query(context.Background(), query, userId)
	if err != nil {
		return BalanceUser{}, err
	}

	users, err := pgx.CollectOneRow[BalanceUser](rows, pgx.RowToStructByName)
	if err != nil {
		return BalanceUser{}, err
	}
	return users, err
}
