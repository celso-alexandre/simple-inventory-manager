package models

import (
	"database/sql"

	"github.com/celso-alexandre/simple-inventory-manager/server/db"
	"github.com/celso-alexandre/simple-inventory-manager/server/utils"
)

type User struct {
	Id              int64  `json:"id"`
	Username        string `json:"username" binding:"required"`
	Password        string `json:"password" binding:"required"`
	UpdatedByUserId int64  `json:"updatedByUserId"`
}

func upsertUser(u *User, sql string) (*sql.Row, error) {
	passwdHash, err := utils.HashPassword(u.Password)
	if err != nil {
		return nil, err
	}
	return db.DB.QueryRow(sql, u.Username, passwdHash, u.UpdatedByUserId), nil
}

func (u *User) Create() error {
	const sql = `
		INSERT INTO "Users" ("username", "password", "updatedByUserId") 
		VALUES ($1, $2, $3)
		RETURNING "id", "password", "updatedByUserId"
	`
	result, err := upsertUser(u, sql)
	if err != nil {
		// fmt.Println("Error upsertUser 01")
		return err
	}
	err = result.Scan(&u.Id, &u.Password, &u.UpdatedByUserId)
	if err != nil {
		// fmt.Println("Error upsertUser 02")
		return err
	}
	return nil
}

func (u *User) UpdatePassword() error {
	const sql = `
		UPDATE "Users" 
		SET    "password" = $2,
				 "updatedAt" = now(),
				 "updatedByUserId" = CASE WHEN $3 <= 0 THEN "Users"."id" ELSE $3 END
		WHERE  "username" = $1
		RETURNING "id"
	`
	result, err := upsertUser(u, sql)
	if err != nil {
		// fmt.Println("Error updatePassword 01")
		return err
	}
	err = result.Scan(&u.Id)
	if err != nil {
		// fmt.Println("Error updatePassword 02")
		return err
	}
	return nil
}

func FindUserByUsername(username string) (*User, error) {
	const sql = `
		SELECT id, username, coalesce(password, '') as password
		FROM "Users"
		WHERE username = $1
	`
	// fmt.Println("username: ", username)
	// fmt.Println("sql: ", sql)
	// fmt.Println("Before Before Scanned")
	row := db.DB.QueryRow(sql, username)
	// fmt.Println("Before Scanned")
	u := User{}
	err := row.Scan(&u.Id, &u.Username, &u.Password)
	return &u, err
}

// func FindUserById(id int64) (*User, error) {
// 	const sql = `
// 		SELECT id, username, password
// 		FROM users
// 		WHERE id = $1
// 	`
// 	row := db.DB.QueryRow(sql, id)
// 	var u User
// 	err := row.Scan(&u.Id, &u.Username, &u.Password)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &u, nil
// }

// func (u *User) Update() error {
// 	const sql = `
// 		UPDATE users
// 		SET username = $1, password = $2
// 		WHERE id = $3
// 	`
// 	_, err := db.DB.Exec(sql, u.Username, u.Password, u.Id)
// 	return err
// }

// func FindAllUsers() ([]User, error) {
// 	const sql = `SELECT id, username, password FROM users`
// 	rows, err := db.DB.Query(sql)
// 	if err != nil {
// 		return nil, err
// 	}
// 	defer rows.Close()
// 	users := []User{}
// 	for rows.Next() {
// 		var u User
// 		err := rows.Scan(&u.Id, &u.Username, &u.Password)
// 		if err != nil {
// 			return nil, err
// 		}
// 		u.Password = ""
// 		users = append(users, u)
// 	}
// 	return users, nil
// }

// func DeleteUser(id int64) error {
// 	const sql = `DELETE FROM users WHERE id = $1`
// 	_, err := db.DB.Exec(sql, id)
// 	return err
// }
