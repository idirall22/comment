package comment

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/idirall22/utilities"
)

// AddComment handler
func (s *Service) AddComment(w http.ResponseWriter, r *http.Request) {

	userID, err := utilities.GetUserIDFromContext(r.Context())

	if err != nil {

		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	form := CForm{}
	err = json.NewDecoder(r.Body).Decode(&form)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Could not add your comment", http.StatusBadRequest)
		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	c, err := s.addComment(ctx, userID, form)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, "Error Serveur", http.StatusInternalServerError)
		return
	}
}

// UpdateComment handler
func (s *Service) UpdateComment(w http.ResponseWriter, r *http.Request) {

	userID, err := utilities.GetUserIDFromContext(r.Context())

	if err != nil {

		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	commentID, err := utilities.GetURLID(r, "")

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	err = utilities.ValidateID(commentID)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	form := CForm{}

	err = json.NewDecoder(r.Body).Decode(&form)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Could not add your comment", http.StatusBadRequest)
		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	c, err := s.updateComment(ctx, userID, commentID, form)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)

	if err := json.NewEncoder(w).Encode(c); err != nil {
		http.Error(w, "Error Serveur", http.StatusInternalServerError)
		return
	}
}

// DeleteComment handler
func (s *Service) DeleteComment(w http.ResponseWriter, r *http.Request) {

	userID, err := utilities.GetUserIDFromContext(r.Context())

	if err != nil {

		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	commentID, err := utilities.GetURLID(r, "")

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	err = utilities.ValidateID(commentID)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	err = s.deleteComment(ctx, userID, commentID)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
