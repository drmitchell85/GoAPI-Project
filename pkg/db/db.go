package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "marty"
	dbname   = "projectonedb"
)

func DbConnection() *sql.DB {

	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("Tried to open connect to DB. Err: ", err)
	}

	log.Println("Checking if db is present...")
	if err = db.Ping(); err != nil {
		log.Println("Database not present, building...")
		insertDB()
	} else {
		log.Println("Database already present")
	}

	insertTables(db)

	fmt.Println("Connected to db")
	return db
}

func insertDB() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable", host, port, user, password)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalln("Tried to open connect to DB. Err: ", err)
	}

	// insert db using query
	txt := "CREATE DATABASE " + dbname
	_, err = db.Exec(txt)
	if err != nil {
		log.Fatalln("Add DB query failed. Err: ", err)
	}

	log.Println("Created db")
}

func insertTables(db *sql.DB) {
	// get sql file and read
	query, err := ioutil.ReadFile("../pkg/db/create_tables.sql")
	if err != nil {
		log.Fatalln("Could not read file. Err: ", err)
	}

	_, err = db.Exec(string(query))
	if err != nil {
		log.Fatalln("Add tables query failed. Err: ", err)
	}
}
