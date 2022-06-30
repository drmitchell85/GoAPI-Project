package apis

import (
	"database/sql"
	"log"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var art = &Article{
	Id:      "art_1",
	Title:   "Title of the article",
	Desc:    "Desc of the article",
	Content: "Here is article content",
}

func NewMock() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub db conn", err)
	}

	return db, mock
}

func TestGetArticleByID(t *testing.T) {
	db, mock := NewMock()

	qryTxt := `SELECT * FROM article WHERE article_id = $1`

	rows := sqlmock.NewRows([]string{"article_id", "articleTitle", "articleDesc", "articleContent"}).
		AddRow(art.Id, art.Title, art.Desc, art.Title)

	mock.ExpectQuery(regexp.QuoteMeta(qryTxt)).WithArgs(art.Id).WillReturnRows(rows)

	article, err := getArticleByID(art.Id, db)
	assert.NotNil(t, article)
	assert.NoError(t, err)
}
