package provider

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

var (
	// ErrorForeignKey when there is not a reference
	ErrorForeignKey = errors.New("Foreign Key violation")
	// ErrorNoRow when there comment not exists
	ErrorNoRow = errors.New("comment not exists")
	// ErrorUnique when there is already a comment with same id
	ErrorUnique = errors.New("There is already a comment with this id")
	// ErrorServer when there is an error
	ErrorServer = errors.New("Error server")
)

func parseError(err error) error {
	if err == sql.ErrNoRows {
		return ErrorNoRow
	}
	if e, ok := err.(*pq.Error); ok {

		switch e.Code.Name() {
		case "23503":
			return ErrorForeignKey
		case "23505":
			return ErrorUnique
		}
	}

	return err
}
