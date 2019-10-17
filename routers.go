package comment

import "github.com/gorilla/mux"

// Router comment endpoints
func (s *Service) Router() *mux.Router {
	r := &mux.Router{}

	r.HandleFunc("/add", s.AddComment).Methods("POST")
	r.HandleFunc("/update/{id}", s.UpdateComment).Methods("POST")
	r.HandleFunc("/delete/{id}", s.DeleteComment).Methods("POST")
	return r
}
