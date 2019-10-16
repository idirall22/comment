package comment

import (
	"encoding/json"
	"net/http"
)

// AddComment handler
func AddComment(w http.ResponseWriter, r *http.Request) {

	form := CForm{}
	err := json.NewDecoder(r.Body).Decode(&form)

	defer r.Body.Close()

	if err != nil {
		return
	}

	if !form.ValidateForm() {
		return
	}

}

// UpdateComment handler
func UpdateComment(w http.ResponseWriter, r *http.Request) {

}

// DeleteComment handler
func DeleteComment(w http.ResponseWriter, r *http.Request) {

}
