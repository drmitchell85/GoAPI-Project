package apis

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	errorhandlers "goapiproject.com/pkg/error"
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
	errorhandlers.CheckError(err)

	fmt.Fprintf(w, string(out))
}
