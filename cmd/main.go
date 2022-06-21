package main

import (
	"goapiproject.com/pkg/apis"
	"goapiproject.com/pkg/db"

	_ "github.com/lib/pq"
)

func main() {
	db.DbConnection()
	apis.HandleRequests()
}
