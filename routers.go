package comment

import (
	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
)

// Router comment endpoints
func (s *Service) Router(r *mux.Router) {

	sr := r.PathPrefix("/comments").Subrouter()

	sr.HandleFunc("/", u.AuthnticateUser(s.AddComment)).Methods("POST")
	sr.HandleFunc("/{id}", u.AuthnticateUser(s.UpdateComment)).Methods("PUT")
	sr.HandleFunc("/{id}", u.AuthnticateUser(s.DeleteComment)).Methods("DELETE")
	sr.HandleFunc("/stream", u.AuthnticateUser(s.SubscribeCommentStream))

}
