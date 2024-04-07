package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB


func InitDB() {
	
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")

	if err != nil {
		panic("Error loading .env file")
	}
	dbURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPassword, dbHost, dbPort, dbName)

	// DB, err = sql.Open("mysql", "root:saurabh@tcp(localhost:3306)/rest_api_go?parseTime=true")
	DB, err = sql.Open("mysql", dbURI)

	if err != nil {
		panic(err)
	}
	// defer DB.Close()

	createTables(DB)

}

func createTables(DB *sql.DB) {
	createUsersTable := `CREATE TABLE IF NOT EXISTS users (id INTEGER AUTO_INCREMENT PRIMARY KEY, email VARCHAR(255) NOT NULL UNIQUE,password TEXT NOT NULL)`
	_, err := DB.Exec(createUsersTable)
	if err != nil {
		panic(err.Error())
	}
	createEventsTable := `CREATE TABLE IF NOT EXISTS events(
		id INTEGER AUTO_INCREMENT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		dateTime TIMESTAMP NOT NULL,
		user_id INTEGER,
		FOREIGN KEY (user_id) REFERENCES users(id)

	)`
	_, err = DB.Exec(createEventsTable)

	if err != nil {
		panic(err.Error())
	}
	createRegistrationsTable := `CREATE TABLE IF NOT EXISTS registrations(
		id INTEGER AUTO_INCREMENT PRIMARY KEY,
		event_id INTEGER,
		user_id INTEGER,
		FOREIGN KEY (event_id) REFERENCES events(id),
		FOREIGN KEY (user_id) REFERENCES users(id)

	)`
	_, err = DB.Exec(createRegistrationsTable)

	if err != nil {
		panic(err.Error())
	}

}
