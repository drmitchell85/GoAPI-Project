package apis

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
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

func TestHealth(t *testing.T) {
	// Create a request to pass to our handler. We don't have any query parameters for now, so
	// we'll pass 'nil' as the third parameter.
	req, err := http.NewRequest("GET", "/health-check", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	db, _ := NewMock()
	handlers := dbStruct{db: db}
	handler := http.HandlerFunc(handlers.health)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check for expected status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned incorrect status code. expected %v, received %v",
			http.StatusOK, status)
	}

	// check for expected response body
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: expected %v, received %v",
			expected, rr.Body.String())
	}
}

// figure out how to test DB connections as well
func TestGetAllArticlesAPI(t *testing.T) {
	db, mock := NewMock()
	rows := sqlmock.NewRows([]string{"article_id", "articleTitle", "articleDesc", "articleContent"}).
		AddRow(as[0].Id, as[0].Title, as[0].Desc, as[0].Title).
		AddRow(as[1].Id, as[1].Title, as[1].Desc, as[1].Title)
	qryTxt := `SELECT * FROM article`
	mock.ExpectQuery(regexp.QuoteMeta(qryTxt)).WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/articles/list", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handlers := dbStruct{db}
	handler := http.HandlerFunc(handlers.getAllArticlesAPI)
	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("expected: %v, received: %v", http.StatusOK, status)
	}

	strBody := rr.Body.String()
	response := ArticleResponse{}
	json.Unmarshal([]byte(strBody), &response)

	assert.EqualValues(t, 2, response.Count)
	assert.ElementsMatch(t, as, response.Articles)
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
	assert.NotNil(t, scan)
	assert.NoError(t, err)
	assert.Equal(t, a1.Id, scan[0].Id)
}
