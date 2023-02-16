package database

import (
	"log"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

// register a user
func User_register(username string, password string, status uint) int64 {
	stmt, err := DB.Prepare("INSERT INTO users(uuid, username, password, status) values (?,?,?,?);")
	if err != nil {
		log.Fatal(err)
	}

	uuid := uuid.New()
	res, err := stmt.Exec(uuid.String(), username, password, status)
	if err != nil {
		log.Fatal(err)
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return lid
}

// delete a user
func User_delete_by_uuid(uuid string) int64 {
	stmt, err := DB.Prepare("DELETE FROM users where uuid=?;")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(uuid)
	if err != nil {
		log.Fatal(err)
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return lid
}

// delete a user
func User_delete_by_username(username string) int64 {
	stmt, err := DB.Prepare("DELETE FROM users where username=?;")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(username)
	if err != nil {
		log.Fatal(err)
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return lid
}

// change password
func User_change_password_by_username(username string, password string) int64 {
	stmt, err := DB.Prepare("UPDATE users SET password=? where username=?;")
	if err != nil {
		log.Fatal(err)
	}

	res, err := stmt.Exec(password, username)
	if err != nil {
		log.Fatal(err)
	}

	lid, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	return lid
}

// users
type User struct {
	Uuid     string
	Username string
	Password string
	Status   uint
}

func User_by_username(username string) User {
	stmt, err := DB.Prepare("SELECT * FROM users WHERE username=?;")
	if err != nil {
		log.Fatal(err)
	}

	row, err := stmt.Query(username)
	if err != nil {
		log.Fatal(err)
	}

	var u User

	row.Next()
	row.Scan(&u.Uuid, &u.Username, &u.Password, &u.Status)

	return u
}
