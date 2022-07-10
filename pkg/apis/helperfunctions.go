package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Id      string `json:"article_id"`
	Title   string `json:"articleTitle"`
	Desc    string `json:"articleDesc"`
	Content string `json:"articleContent"`
}

type ArticleResponse struct {
	Status   string
	Code     int
	Count    int
	Articles []Article
}

func getAllArticles(db *sql.DB) (*sql.Rows, error) {
	qryTxt := `SELECT * FROM article`
	rows, err := db.Query(qryTxt)
	return rows, err
}

func getArticleByID(article_id string, db *sql.DB) (*sql.Rows, error) {
	qryTxt := `SELECT * FROM article WHERE article_id = $1`
	result, err := db.Query(qryTxt, article_id)
	return result, err
}

func insertArticle(a Article, db *sql.DB) (*sql.Rows, error) {
	qryTxt := `INSERT into "article"(articletitle, articledesc, articlecontent) VALUES($1, $2, $3) RETURNING *`
	result, err := db.Query(qryTxt, a.Title, a.Desc, a.Content)
	fmt.Println("err:", err)
	return result, err
}

func scanArticles(rows *sql.Rows) ([]Article, error) {
	var articles []Article

	for rows.Next() {
		var article Article
		if err := rows.Scan(&article.Id, &article.Title, &article.Desc, &article.Content); err != nil {
			return articles, err
		}
		articles = append(articles, article)
	}

	return articles, nil
}

func articlesResponse(status string, code int, articles []Article, w http.ResponseWriter) {
	response := ArticleResponse{
		Status:   status,
		Code:     code,
		Count:    len(articles),
		Articles: articles,
	}

	out, err := json.Marshal(response)
	if err != nil {
		articlesErrResponse("error encoding response: %s", err, articles, w)
		return
	}

	fmt.Fprintln(w, string(out))
}

func articlesErrResponse(msg string, err error, articles []Article, w http.ResponseWriter) {
	msgFmt := fmt.Sprintf(msg, err)
	log.Println(msgFmt)
	articlesResponse(msg, 400, articles, w)
}
