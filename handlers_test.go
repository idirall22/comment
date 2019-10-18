package comment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
)

func makeRequestTest(r *http.Request, f http.HandlerFunc) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()
	h := http.HandlerFunc(u.AuthnticateUser(f))
	h.ServeHTTP(w, r)

	return w
}

func testAddCommentHandler(t *testing.T) {

	testData := CForm{Content: "comment test", PostID: 1}
	b, err := json.Marshal(testData)
	if err != nil {
		t.Error(err)
		return
	}

	body := bytes.NewReader(b)

	r, err := http.NewRequest("POST", "/comment", body)
	if err != nil {

	}
	r.Header.Add("Authorization", testToken)
	res := makeRequestTest(r, testService.AddComment)

	if res.Code != http.StatusCreated {
		t.Errorf("Status code should be %d but got %d", res.Code, http.StatusCreated)
	}
}

func testUpdateCommentHandler(t *testing.T) {

	testData := CForm{Content: "comment updated", PostID: 1}

	b, err := json.Marshal(testData)
	if err != nil {
		t.Error(err)
		return
	}

	body := bytes.NewReader(b)

	res := httptest.NewRecorder()

	r, err := http.NewRequest("PUT", "/comment/1", body)
	if err != nil {

	}

	r.Header.Add("Authorization", testToken)

	mux.SetURLVars(r, map[string]string{"id": "1"})

	router := mux.NewRouter()
	router.HandleFunc("/comment/{id}", u.AuthnticateUser(testService.UpdateComment)).Methods("PUT")
	router.ServeHTTP(res, r)

	if res.Code != http.StatusNoContent {
		t.Errorf("Status code should be %d but got %d", res.Code, http.StatusNoContent)
	}
}

func testDeleteCommentHandler(t *testing.T) {

	res := httptest.NewRecorder()
	r, err := http.NewRequest("DELETE", "/comment/1", nil)

	if err != nil {
		t.Error(err)
		return
	}

	r.Header.Add("Authorization", testToken)

	mux.SetURLVars(r, map[string]string{"id": "1"})

	router := mux.NewRouter()
	router.HandleFunc("/comment/{id}", u.AuthnticateUser(testService.DeleteComment)).Methods("DELETE")
	router.ServeHTTP(res, r)

	if res.Code != http.StatusNoContent {
		t.Errorf("Status code should be %d but got %d", res.Code, http.StatusNoContent)
	}
}
