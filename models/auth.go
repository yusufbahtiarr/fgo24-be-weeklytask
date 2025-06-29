package models

import (
	"context"
	"fgo24-be-ewallet/utils"

	"github.com/jackc/pgx/v5/pgconn"
)

func RegisterUser(user User) error {
	conn, err := utils.DBConnect()
	if err != nil {
		return err
	}
	defer conn.Close()

	_, err = conn.Exec(
		context.Background(),
		`INSERT INTO users (email, password, pin) VALUES ($1, $2, $3)`,
		user.Email,
		user.Password,
		user.Pin,
	)

	if err != nil {
		if pgErr, ok := err.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			return err
		}
		return err
	}

	return nil
}
