package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type dbStruct struct {
	db *sql.DB
}

func (d dbStruct) health(w http.ResponseWriter, req *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	io.WriteString(w, `{"alive": true}`)
}

func (d dbStruct) getArticleByIdAPI(w http.ResponseWriter, req *http.Request) {
	urlArray := strings.Split(req.URL.String(), "=")
	article_id := urlArray[len(urlArray)-1]
	var articles []Article

	result, err := getArticleByID(article_id, d.db)
	if err != nil {
		articlesErrResponse("error getting article by ID: %s", err, articles, w)
		return
	}

	articles, err = scanArticles(result)
	if err != nil {
		articlesErrResponse("error scanning articles: %s", err, articles, w)
		return
	}

	if len(articles) == 0 {
		msg := fmt.Sprintf("Failure article with ID %s does not exist", article_id)
		articlesResponse(msg, 200, articles, w)
		return
	}

	articlesResponse("Success", 200, articles, w)
}

func (d dbStruct) getAllArticlesAPI(w http.ResponseWriter, req *http.Request) {
	var articles []Article
	rows, err := getAllArticles(d.db)
	if err != nil {
		articlesErrResponse("error retrieving aricles: %s", err, articles, w)
		return
	}
	defer rows.Close()

	articles, err = scanArticles(rows)
	if err != nil {
		articlesErrResponse("error scanning articles: %s", err, articles, w)
		return
	}

	articlesResponse("Success", 200, articles, w)
}

func (d dbStruct) insertArticleAPI(w http.ResponseWriter, req *http.Request) {
	var a Article
	var articles []Article
	err := json.NewDecoder(req.Body).Decode(&a)
	if err != nil {
		articlesErrResponse("error decoding body: %s", err, articles, w)
		return
	}

	result, err := insertArticle(a, d.db)
	if err != nil {
		articlesErrResponse("error inserting article: %s", err, articles, w)
		return
	}

	articles, err = scanArticles(result)
	if err != nil {
		articlesErrResponse("error scanning articles: %s", err, articles, w)
		return
	}

	articlesResponse("Success", 200, articles, w)
}
