package apis

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"goapiproject.com/pkg/db"
	errorhandlers "goapiproject.com/pkg/error"
)

func returnArticle(w http.ResponseWriter, req *http.Request) {
	fmt.Println("running returnArticle...")
	urlArray := strings.Split(req.URL.String(), "=")
	article_id := urlArray[len(urlArray)-1]

	qryTxt := `SELECT * FROM article WHERE article_id = $1`

	fmt.Println("running query...")

	result, err := db.PsqlDB.Query(qryTxt, article_id)
	fmt.Println("query err: ", err)
	errorhandlers.CheckError(err)

	fmt.Println("scanning articles...")

	articles, error := scanArticles(result)
	errorhandlers.CheckError(error)

	if len(articles) == 0 {
		msg := "Failure article with ID does not exist"
		enf := errorhandlers.ErrNotFound{
			Url:     req.URL.String(),
			Code:    400,
			Message: msg,
		}
		log.Println("enf: ", enf)

		articlesResponse(msg, 400, articles, w)
		return
	}

	articlesResponse("Success", 200, articles, w)
}

func returnAllArticles(w http.ResponseWriter, req *http.Request) {
	qryTxt := `SELECT * FROM article`
	rows, err := db.PsqlDB.Query(qryTxt)
	errorhandlers.CheckError(err)
	defer rows.Close()

	articles, error := scanArticles(rows)
	errorhandlers.CheckError(error)

	articlesResponse("Success", 200, articles, w)
}

func insertArticle(w http.ResponseWriter, req *http.Request) {
	var a Article
	err := json.NewDecoder(req.Body).Decode(&a)
	errorhandlers.CheckError(err)

	qryTxt := `insert into "article"(title, article_description, content) values($1, $2, $3) RETURNING *`
	result, err2 := db.PsqlDB.Query(qryTxt, a.Title, a.Desc, a.Content)
	errorhandlers.CheckError(err2)

	articles, error := scanArticles(result)
	errorhandlers.CheckError(error)

	articlesResponse("Success", 200, articles, w)
}
