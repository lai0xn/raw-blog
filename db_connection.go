package main

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Connect establishes the connection to the database
func Connect() error {
	db, err := sql.Open("sqlite3", "./database.sqlite3")
	if err != nil {
		return err
	}
	log.Printf("%s Sqlite Database Opened", time.Now().String())

	// create the articles table if it doesn't already exist

	sqlistm := `
  create table if not exists articles (id integer not null primary key autoincrement, title text, content text);
  `
	_, err = db.Exec(sqlistm)
	if err != nil {
		return err
	}
	DB = db
	return nil
}
