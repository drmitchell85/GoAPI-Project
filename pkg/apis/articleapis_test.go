package apis

import (
	"database/sql"
	"fmt"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var a1 = &Article{
	Id:      "art_1",
	Title:   "Title of the article",
	Desc:    "Desc of the article",
	Content: "Here is article content",
}

var a2 = &Article{
	Id:      "art_2",
	Title:   "Title of the article2",
	Desc:    "Desc of the article2",
	Content: "Here is article2 content",
}

var as = []Article{
	{a1.Id, a1.Title, a1.Desc, a1.Title},
	{a2.Id, a2.Title, a2.Desc, a2.Title},
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub db conn", err)
	}

	return db, mock
}

func TestGetAllArticles(t *testing.T) {
	db, mock := NewMock()

	rows := sqlmock.NewRows([]string{"article_id", "articleTitle", "articleDesc", "articleContent"}).
		AddRow(as[0].Id, as[0].Title, as[0].Desc, as[0].Title).
		AddRow(as[1].Id, as[1].Title, as[1].Desc, as[1].Title)

	qryTxt := `SELECT * FROM article`
	mock.ExpectQuery(regexp.QuoteMeta(qryTxt)).WillReturnRows(rows)

	articles, err := getAllArticles(db)
	assert.NotNil(t, articles)
	assert.NoError(t, err)

	scan, err := scanArticles(articles)
	assert.NotNil(t, scan)
	assert.NoError(t, err)
	assert.ElementsMatch(t, as, scan)
}

func TestGetArticleByID(t *testing.T) {
	db, mock := NewMock()

	qryTxt := `SELECT * FROM article WHERE article_id = $1`

	rows := sqlmock.NewRows([]string{"article_id", "articleTitle", "articleDesc", "articleContent"}).
		AddRow(a1.Id, a1.Title, a1.Desc, a1.Title)

	mock.ExpectQuery(regexp.QuoteMeta(qryTxt)).WithArgs(a1.Id).WillReturnRows(rows)

	article, err := getArticleByID(a1.Id, db)
	assert.NotNil(t, article)
	assert.NoError(t, err)

	scan, err := scanArticles(article)
	fmt.Println("scan: ", scan)
	assert.NotNil(t, scan)
	assert.NoError(t, err)
	assert.Equal(t, a1.Id, scan[0].Id)
}
