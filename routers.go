package comment

import (
	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
)

// Router comment endpoints
func (s *Service) Router() *mux.Router {
	r := &mux.Router{}

	r.HandleFunc("/comments", u.AuthnticateUser(s.AddComment)).Methods("POST")
	r.HandleFunc("/comments/{id}", u.AuthnticateUser(s.UpdateComment)).Methods("PUT")
	r.HandleFunc("/comments/{id}", u.AuthnticateUser(s.DeleteComment)).Methods("DELETE")
	r.HandleFunc("/comments/stream", u.AuthnticateUser(s.SubscribeCommentStream))

	return r
}
