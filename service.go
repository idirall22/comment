package comment

import (
	"database/sql"
	"time"

	pr "github.com/idirall22/comment/providers/postgres"
)

// TimeoutRequest time to stop a request
var TimeoutRequest = time.Second * 5

// Service structure
type Service struct {
	provider Provider
}

// StartService comment
func StartService(db *sql.DB) *Service {

	p := &pr.PostgresProvider{DB: db, TableName: "comments"}
	return &Service{provider: p}
}
