// The models package contains type definitions for various data objects,
// as well as methods for performing CRUD operations on tables in the
// SQL database.
package models

import (
	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

// User type groups data describing a REST API user
type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

// Saves a User by adding a row to the users table in the database
func (u User) Save() error {
	query := "INSERT INTO users(email, password) VALUES (?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	_, err = stmt.Exec(u.Email, hashedPassword)

	return err
}
