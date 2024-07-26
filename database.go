package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func createDatabase() error {
	db, err := sql.Open("sqlite3", "database/data.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// Create the users table
	_, err = db.Exec(`
                CREATE TABLE IF NOT EXISTS users (
                        id INTEGER PRIMARY KEY AUTOINCREMENT,
                        username TEXT NOT NULL,
                        email TEXT NOT NULL
                );
        `)
	if err != nil {
		return err
	}

	fmt.Println("Database and table created successfully!")
	return nil
}

func initDb() {
	if err := createDatabase(); err != nil {
		fmt.Println(err)
	}
}
