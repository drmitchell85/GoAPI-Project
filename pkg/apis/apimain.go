package apis

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func HandleRequests(db *sql.DB) {
	handlers := dbStruct{db}
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/health", handlers.health)
	http.HandleFunc("/articles/list", handlers.getAllArticlesAPI)
	http.HandleFunc("/articles/fetch", handlers.getArticleByIdAPI)
	http.HandleFunc("/articles/insert", handlers.insertArticleAPI)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func HomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}
