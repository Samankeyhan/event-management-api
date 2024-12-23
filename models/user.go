package models

import (
	"errors"
	"log"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
    ID       int64  `json:"id"`
    Email    string `json:"email" binding:"required"`
    Password string `json:"password" binding:"required"`
}

func (u *User) Save() error {
	query := `INSERT INTO users(email, password) VALUES (?, ?)`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		
		log.Printf("Error preparing query: %v", err)
		return err
	}

	defer stmt.Close()
	hashedPass, err := utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPass)

	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}
	userId, err := result.LastInsertId()
	if err != nil {
        log.Printf("Error getting last insert ID: %v", err)
        return err
    }

	u.ID = userId
	return nil
}

func (u *User) ValidateUser() error {

	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return err
	}

	passwordIsValid := utils.CheckPassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials Inavlid")
	}

	return nil

}