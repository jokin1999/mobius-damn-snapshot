package database

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

func Init() {
	log.Println("Creating database")

	var err error

	DB, err = sql.Open("sqlite3", "./damn-snapshot.db")
	if err != nil {
		log.Fatal(err)
	}

	// create database table
	sql_table := `
	CREATE TABLE IF NOT EXISTS "users"(
		"uuid" VARCHAR(64) PRIMARY KEY NOT NULL,
		"username" VARCHAR(64) UNIQUE NOT NULL,
		"password" VARCHAR(64) NOT NULL,
		"status"  INT UNSIGNED NOT NULL
	);
	CREATE TABLE IF NOT EXISTS "permissions"(
		"user_uuid" VARCHAR(64) PRIMARY KEY NOT NULL,
		"vmid" TEXT NOT NULL,
		"status"  INT UNSIGNED NOT NULL
	);
	`

	_, err = DB.Exec(sql_table)
	if err != nil {
		log.Fatal(err)
	}

}
