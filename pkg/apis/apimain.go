package apis

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func HandleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/articles/list", returnAllArticles)
	http.HandleFunc("/articles/fetch", returnArticle)
	http.HandleFunc("/articles/insert", insertArticle)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func HomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}
