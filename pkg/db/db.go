package db

import (
	"database/sql"
	"fmt"

	errorhandlers "goapiproject.com/pkg/error"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = ""
	dbname   = "projectonedb"
)

var PsqlDB *sql.DB

func DbConnection() {
	// checkDB()
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	errorhandlers.CheckError(err)

	// close db
	// defer db.Close()

	// check db
	err = db.Ping()
	errorhandlers.CheckError(err)

	fmt.Println("Connected to db")
	PsqlDB = db
}

func checkDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", connStr)
	errorhandlers.CheckError(err)

	// close db
	// defer db.Close()

	// check db
	err = db.Ping()
	errorhandlers.CheckError(err)

	// TODO: get sql file
	// sql.Query()

	fmt.Println("Created db")
	db.Close()
}
