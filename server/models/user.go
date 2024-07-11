package models

import (
	"github.com/igor570/eventexplorer/db"
	"github.com/igor570/eventexplorer/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
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
