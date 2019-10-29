package comment

import (
	"database/sql"

	pr "github.com/idirall22/comment/providers/postgres"
)

// Service structure
type Service struct {
	provider Provider
	broker   Broker
}

// StartService comment
func StartService(db *sql.DB, tableName string) *Service {

	p := &pr.PostgresProvider{DB: db, TableName: tableName}
	return &Service{provider: p}
}
