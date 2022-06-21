package apis

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestReturnArticle(t *testing.T) {

	// Create a request to pass to our handler. We don't have any query parameters for now, so we'll
	// pass 'nil' as the third parameter
	req, err := http.NewRequest("GET", "/articles/fetch", nil)
	if err != nil {
		t.Fatal(err)
	}

	// We create a ResponseRecorder (which satisfies http.ResponseWriter) to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(returnArticle)

	// Our handlers satisfy http.Handler, so we can call their ServeHTTP method
	// directly and pass in our Request and ResponseRecorder.
	handler.ServeHTTP(rr, req)

	// Check the expected status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: received %v, expected %v", status, http.StatusOK)
	}

	// Check the expected response body
	expected := `{"alive": true}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}
