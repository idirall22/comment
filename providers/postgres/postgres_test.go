package provider

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	"github.com/idirall22/utilities"
	_ "github.com/lib/pq"
)

var (
	provider   *PostgresProvider
	database   *sql.DB
	testToken  string
	commentNum = 5
	tableName  = "comments"
	query      = fmt.Sprintf(`
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

	provider = &PostgresProvider{DB: db, TableName: tableName}

	testToken = utilities.LoginUser(db)

	t.Run("New", testNew)
	t.Run("List", testList)
	t.Run("update", testUpdate)
	t.Run("delete", testDelete)
}

// Test New
func testNew(t *testing.T) {
	for i := 0; i < commentNum; i++ {
		_, err := provider.New(context.Background(), fmt.Sprintf("message %v", i), 1, 1)

		if err != nil {
			t.Error("Error should be nil but got:", err)
		}
	}
}

// Test List
func testList(t *testing.T) {
	comments, err := provider.List(context.Background(), 1, 5, 0)

	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}

	if len(comments) != commentNum {
		t.Errorf("Error comments slice length should be %d But got %d",
			commentNum, len(comments))
	}
}

// Test update
func testUpdate(t *testing.T) {
	_, err := provider.Update(context.Background(), 1, 1, "updated message")

	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
}

// Test delete
func testDelete(t *testing.T) {
	err := provider.Delete(context.Background(), 1, 1)
	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
}
