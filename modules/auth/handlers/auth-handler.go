package handlers

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/evan_nurandiz/go_fiber_boilerplate/database"
	"github.com/evan_nurandiz/go_fiber_boilerplate/helpers"
	"github.com/evan_nurandiz/go_fiber_boilerplate/modules/auth/model"
)

var (
	db *sql.DB = database.SetupDatabaseConnection()
)

type AuthHandler interface {
	RegisterUser(model.User)
}

func RegisterUser(user model.User) (string, error) {
	user.Password = helpers.HashAndSalt([]byte(user.Password))
	tx, err := db.BeginTx(context.Background(), nil)
	fmt.Println(user)

	if err != nil {
		return "", err
	}

	_, execErr := tx.Exec(`insert into public.user (email, name, password, created_on) 
	values  ($1 , $2, $3, $4)`,
		user.Email, user.Name, user.Password, time.Now())

	if execErr != nil {
		_ = tx.Rollback()
		return "", execErr
	}

	if err := tx.Commit(); err != nil {
		return "", err
	}

	return user.Email, nil
}

func GetUserDataByEmail(email string) (model.User, string) {
	var user model.User

	row := db.QueryRow(`
		select u.user_id, u.name, u.email, u.password 
		from public.user u
		where u.email = $1
	`, email)

	err := row.Scan(&user.User_id, &user.Name, &user.Email, &user.Password)

	if err != nil {
		fmt.Println(err)
		return model.User{}, "user not found"
	}

	return user, ""
}
