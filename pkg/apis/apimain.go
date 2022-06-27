package apis

import (
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func HandleRequests() {
	http.HandleFunc("/", HomePage)
	http.HandleFunc("/articles/list", getAllArticlesAPI)
	http.HandleFunc("/articles/fetch", getArticleByIdAPI)
	http.HandleFunc("/articles/insert", insertArticleAPI)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func HomePage(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Welcome to the Homepage!")
	fmt.Println("Endpoint Hit: homepage")
}
