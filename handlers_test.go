package comment

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func makeRequestTest(r *http.Request, f http.HandlerFunc) *httptest.ResponseRecorder {

	w := httptest.NewRecorder()
	h := http.HandlerFunc(f)
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

	r, err := http.NewRequest("POST", "/add_comment", body)
	if err != nil {

	}
	res := makeRequestTest(r, testService.AddComment)

	if res.Code != http.StatusCreated {
		t.Errorf("Status code should be %d but got %d", res.Code, http.StatusCreated)
	}
}

func testUpdateCommentHandler(t *testing.T) {
	testData := UForm{Content: "comment test", ID: 1}

	b, err := json.Marshal(testData)
	if err != nil {
		t.Error(err)
		return
	}

	body := bytes.NewReader(b)

	r, err := http.NewRequest("POST", "/update_comment", body)
	if err != nil {

	}

	res := makeRequestTest(r, testService.UpdateComment)

	if res.Code != http.StatusNoContent {
		t.Errorf("Status code should be %d but got %d", res.Code, http.StatusNoContent)
	}
}

func testDeleteCommentHandler(t *testing.T) {

	r, err := http.NewRequest("POST", "/delete/1", nil)

	if err != nil {
		t.Error(err)
		return
	}
	mux.SetURLVars(r, map[string]string{"id": "1"})

	res := httptest.NewRecorder()

	router := mux.NewRouter()
	router.HandleFunc("/delete/{id}", testService.DeleteComment).Methods("POST")
	router.ServeHTTP(res, r)

	if res.Code != http.StatusNoContent {
		t.Errorf("Status code should be %d but got %d", res.Code, http.StatusNoContent)
	}
}
