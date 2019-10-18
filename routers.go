package comment

import (
	"github.com/gorilla/mux"
	u "github.com/idirall22/user"
)

// Router comment endpoints
func (s *Service) Router() *mux.Router {
	r := &mux.Router{}

	r.HandleFunc("/comment", u.AuthnticateUser(s.AddComment)).Methods("POST")
	r.HandleFunc("/comment/{id}", u.AuthnticateUser(s.UpdateComment)).Methods("PUT")
	r.HandleFunc("/comment/{id}", u.AuthnticateUser(s.DeleteComment)).Methods("DELETE")

	return r
}
