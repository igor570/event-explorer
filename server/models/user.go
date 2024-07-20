package models

import (
	"errors"

	"github.com/igor570/eventexplorer/db"
	"github.com/igor570/eventexplorer/utils"
)

type User struct {
	ID       int64
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {

	// Send data to database
	query := `
	INSERT INTO users(email, password)
	VALUES(?, ?)
	`
	val, err := db.Database.Prepare(query)

	if err != nil {
		return err
	}

	defer val.Close() //defer execution until end

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	result, err := val.Exec(u.Email, hashedPassword)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId() //get the last inserted id in DB

	if err != nil {
		return err
	}

	u.ID = id //set the structs id to last id in DB

	return err

}

func (u *User) ValidateCredentials() error {
	query := `SELECT id, password FROM users WHERE email = ?`
	row := db.Database.QueryRow(query, u.Email) //find the row of specific email

	var retrievedPassword string //storing the password sent from req.body
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	isPasswordValid := utils.ValidatePassword(u.Password, retrievedPassword) //compare structs password to retreived password

	if !isPasswordValid {
		return errors.New("invalid credentials")
	}

	return nil
}
