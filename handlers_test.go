package comment

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
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

// Test addCommentHandler
func testAddCommentHandler(t *testing.T) {
	var body io.Reader

	for i := 0; i < 3; i++ {

		switch i {

		case 0:
			// When data is valid
			testData := CForm{Content: "comment test", PostID: 1}
			b, err := json.Marshal(testData)
			if err != nil {
				t.Error(err)
				return
			}
			body = bytes.NewReader(b)
			break

		case 1:
			// When data is not valid
			body = bytes.NewReader(nil)
			break
		case 2:
			// When data is not valid
			testData := CForm{Content: "", PostID: 0}
			b, err := json.Marshal(testData)
			if err != nil {
				t.Error(err)
				return
			}
			body = bytes.NewReader(b)
			break
		}

		r, err := http.NewRequest("POST", "/comment", body)
		if err != nil {
			t.Error("Error to make a new request:", err)
			return
		}

		r.Header.Add("Authorization", testToken)
		res := makeRequestTest(r, testService.AddComment)

		switch i {

		case 0:
			if res.Code != http.StatusCreated {
				t.Errorf("Status code should be %d but got %d", http.StatusCreated, res.Code)
			}

		case 1, 2:
			if res.Code != http.StatusBadRequest {
				t.Errorf("Status code should be %d but got %d", http.StatusBadRequest, res.Code)
			}
		}
	}
}

// Test updateCommentHandler
func testUpdateCommentHandler(t *testing.T) {

	var body io.Reader
	var id int64
	var token string

	for i := 0; i < 4; i++ {

		switch i {

		case 0:
			// When data is valid and comment id are valid
			testData := CForm{Content: "comment updated", PostID: 1}

			b, err := json.Marshal(testData)
			if err != nil {
				t.Error(err)
				return
			}
			body = bytes.NewReader(b)
			id = 1
			token = testToken

		case 1:
			// When data is valid but comment id is not valid
			testData := CForm{Content: "comment updated", PostID: 1}

			b, err := json.Marshal(testData)
			if err != nil {
				t.Error(err)
				return
			}
			body = bytes.NewReader(b)
			id = 0
			token = testToken

		case 2:
			// When data is not valid but comment id is valid
			body = bytes.NewReader(nil)
			id = 1
			token = testToken
		case 3:
			// When user is not authorized
			body = bytes.NewReader(nil)
			id = 1
			token = ""
		}

		res := httptest.NewRecorder()

		r, err := http.NewRequest("PUT", fmt.Sprintf("/comment/%v", id), body)
		if err != nil {

		}

		r.Header.Add("Authorization", token)

		mux.SetURLVars(r, map[string]string{"id": "1"})

		router := mux.NewRouter()
		router.HandleFunc("/comment/{id}", u.AuthnticateUser(testService.UpdateComment)).Methods("PUT")
		router.ServeHTTP(res, r)

		switch i {

		case 0:
			if res.Code != http.StatusNoContent {
				t.Errorf("Status code should be %d but got %d", http.StatusNoContent, res.Code)
			}

		case 1, 2:
			if res.Code != http.StatusBadRequest {
				t.Errorf("Status code should be %d but got %d", http.StatusBadRequest, res.Code)
			}

		case 3:
			if res.Code != http.StatusUnauthorized {
				t.Errorf("Status code should be %d but got %d", http.StatusUnauthorized, res.Code)
			}
		}
	}

}

// Test deleteCommentHandler
func testDeleteCommentHandler(t *testing.T) {

	var id int64
	var token string

	for i := 0; i < 3; i++ {

		switch i {

		case 0:
			// When token and id are valid
			token = testToken
			id = 1

		case 1:
			// When token is not valid
			token = ""
			id = 1

		case 2:
			// When id is not valid
			id = 0
			token = testToken
		}

		res := httptest.NewRecorder()
		r, err := http.NewRequest("DELETE", fmt.Sprintf("/comment/%v", id), nil)

		if err != nil {
			t.Error(err)
			return
		}

		r.Header.Add("Authorization", token)

		mux.SetURLVars(r, map[string]string{"id": "1"})

		router := mux.NewRouter()
		router.HandleFunc("/comment/{id}", u.AuthnticateUser(testService.DeleteComment)).Methods("DELETE")
		router.ServeHTTP(res, r)

		switch i {
		case 0:
			if res.Code != http.StatusNoContent {
				t.Errorf("Status code should be %d but got %d", http.StatusNoContent, res.Code)
			}
			break
		case 1:
			if res.Code != http.StatusUnauthorized {
				t.Errorf("Status code should be %d but got %d", http.StatusUnauthorized, res.Code)
			}
			break

		case 2:
			if res.Code != http.StatusBadRequest {
				t.Errorf("Status code should be %d but got %d", http.StatusBadRequest, res.Code)
			}
			break
		}

	}
}
