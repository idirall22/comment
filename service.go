package comment

import (
	"database/sql"

	pr "github.com/idirall22/comment/providers/postgres"
)

// Service structure
type Service struct {
	provider Provider
}

// StartService comment
func StartService(db *sql.DB) *Service {

	p := &pr.PostgresProvider{DB: db, TableName: "comments"}
	return &Service{provider: p}
}
