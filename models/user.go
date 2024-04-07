package models

import (
	"errors"

	"github.com/saurabhpawar8/golang-rest-api/db"
	"github.com/saurabhpawar8/golang-rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err

	}

	defer stmt.Close()
	hashedPassword, err := utils.HashPassword(u.Password)
	if err != nil {
		panic(err)
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return err
	}
	userId, err := result.LastInsertId()

	u.ID = userId
	// var id int64
	// err = stmt.QueryRow(u.Email, hashedPassword).Scan(&id)
	// if err != nil {
	// 	panic(err)
	// }

	// u.ID = id
	return err
}

func (u *User) Validate() error {
	query := `SELECT id, password FROM users WHERE email= ?`
	row := db.DB.QueryRow(query, u.Email)

	var retrivedPassword string
	err := row.Scan(&u.ID, &retrivedPassword)
	if err != nil {
		return err
	}
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrivedPassword)
	if !passwordIsValid {
		return errors.New("Credentials invalid")
	}
	return nil

}
