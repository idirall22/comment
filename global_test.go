package comment

import (
	"database/sql"
	"fmt"
	"testing"

	pr "github.com/idirall22/comment/providers/postgres"
	"github.com/idirall22/utilities"
	_ "github.com/lib/pq"
)

var (
	testService *Service
	database    *sql.DB
	testToken   string
	tableName   = "comments"
	query       = fmt.Sprintf(`
	DROP TABLE IF EXISTS %s;

	CREATE TABLE IF NOT EXISTS %s(
		id SERIAL PRIMARY KEY,
		content VARCHAR NOT NULL,
		user_id INTEGER REFERENCES users(id),
		post_id INTEGER REFERENCES posts(id),
		created_at TIMESTAMP with TIME ZONE DEFAULT now(),
		deleted_at TIMESTAMP DEFAULT NULL
	);
	`, tableName, tableName)
)

// TestGlobal run tests
func TestGlobal(t *testing.T) {

	db, err := utilities.ConnectDataBaseTest()

	if err != nil {
		t.Error(err)
		return
	}

	err = utilities.BuildDataBase(db, query)

	if err != nil {
		t.Error(err)
		return
	}

	defer utilities.CloseDataBaseTest(db)

	provider := &pr.PostgresProvider{DB: db, TableName: tableName}
	testService = &Service{provider: provider}

	testToken = utilities.LoginUser(db)

	t.Run("add a comment handler", testAddCommentHandler)
	t.Run("update a comment handler", testUpdateCommentHandler)
	t.Run("delete a comment handler", testDeleteCommentHandler)
}
