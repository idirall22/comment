package comment

import (
	"context"
	"encoding/json"
	"net/http"
)

// AddComment handler
func (s *Service) AddComment(w http.ResponseWriter, r *http.Request) {

	form := CForm{}
	err := json.NewDecoder(r.Body).Decode(&form)

	defer r.Body.Close()

	if err != nil {
		http.Error(w, "Could not add your comment", http.StatusBadRequest)
		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	c, err := s.addComment(ctx, form)

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

	commentID, err := parseURL(r)
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

	err = s.updateComment(ctx, commentID, form)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}

// DeleteComment handler
func (s *Service) DeleteComment(w http.ResponseWriter, r *http.Request) {

	commentID, err := parseURL(r)
	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	ctx, f := context.WithTimeout(r.Context(), TimeoutRequest)
	defer f()

	err = s.deleteComment(ctx, commentID)

	if err != nil {

		message, code := parseError(err)
		http.Error(w, message, code)

		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
}
