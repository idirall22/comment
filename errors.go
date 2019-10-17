package comment

import (
	"errors"
	"net/http"

	pr "github.com/idirall22/comment/providers/postgres"
)

var (
	// ErrorForm when CForm is not valid
	ErrorForm = errors.New("comment infos not valid")

	// ErrorParam when id is not valid
	ErrorParam = errors.New("ID comment not valid")
)

func parseError(err error) (string, int) {

	message := ""
	code := http.StatusBadRequest

	switch err {

	case ErrorForm:
		message = ErrorForm.Error()
		break

	case ErrorParam:
		message = ErrorParam.Error()
		break

	case pr.ErrorForeignKey:
		message = "Data not valid"

		code = http.StatusConflict
		break

	case pr.ErrorNoRow:
		message = pr.ErrorNoRow.Error()
		code = http.StatusNotFound
		break

	case pr.ErrorUnique:
		message = pr.ErrorUnique.Error()
		code = http.StatusConflict
		break

	case pr.ErrorServer:
		message = pr.ErrorServer.Error()
		code = http.StatusInternalServerError
		break
	}

	return message, code
}
