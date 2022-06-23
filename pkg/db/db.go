package db

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"

	errorhandlers "goapiproject.com/pkg/error"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "marty"
	dbname   = "projectonedb"
)

var PsqlDB *sql.DB

func DbConnection() {
	checkDB()

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

	// check db
	err = db.Ping()
	errorhandlers.CheckError(err)

	// if dir, err := os.Getwd(); err != nil {
	// 	log.Fatalln("Err: ", err)
	// } else {
	// 	log.Println("dir: ", dir)
	// }

	cmd := exec.Command("psql", "-d", dbname, "-f", "../pkg/db/create_tables.sql")
	log.Println("cmd: ", cmd)

	stderr, err := cmd.StderrPipe()
	if err != nil {
		log.Println("err1: ", err)
		panic(err)
	}

	if err := cmd.Start(); err != nil {
		log.Println("err2: ", err)
		panic(err)
	}

	errout, _ := ioutil.ReadAll(stderr)
	if err := cmd.Wait(); err != nil {
		fmt.Println(errout)
		log.Println("err3: ", err)
		panic(err)
	}

	if query, err := ioutil.ReadFile("../pkg/db/create_tables.sql"); err != nil {
		log.Fatalln("Err: ", err)
	} else {
		log.Println("query: ", query)
	}

	// sql.Query()

	fmt.Println("Created db")
	db.Close()
}
