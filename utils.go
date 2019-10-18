package comment

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// parse url to get id
func parseURL(r *http.Request) (int64, error) {

	idString, ok := mux.Vars(r)["id"]

	if !ok {
		return 0, ErrorParam
	}

	id, e := strconv.ParseInt(idString, 10, 64)

	if e != nil {
		return 0, ErrorParam
	}

	if id < 1 {
		return 0, ErrorParam
	}
	return id, nil
}
