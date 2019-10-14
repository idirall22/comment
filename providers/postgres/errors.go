package provider

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

var (
	ErrorForeignKey = errors.New("Error foreign key")
	ErrorNoRow      = errors.New("not exists")
	ErrorUnique     = errors.New("There is already a comment")
	ErrorServer     = errors.New("Error server")
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

	return ErrorServer
}
