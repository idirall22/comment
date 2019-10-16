package provider

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "diskshar_test"
)

const (
	commentNum = 5
)

var provider *PostgresProvider

func cleanDB(db *sql.DB) error {
	query := fmt.Sprintf(`
		DROP TABLE IF EXISTS comments;

		CREATE TABLE IF NOT EXISTS comments(
		    id SERIAL PRIMARY KEY,
			content VARCHAR NOT NULL,
			user_id INTEGER NOT NULL,
			post_id INTEGER NOT NULL,
		    created_at TIMESTAMP with TIME ZONE DEFAULT now(),
		    deleted_at TIMESTAMP DEFAULT NULL
		);
		`)

	_, err := db.Exec(query)

	if err != nil {
		return err
	}
	return nil
}

func closeDB(db *sql.DB) {
	db.Close()
}

func connectDB() error {

	dbInfos := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbInfos)
	if err != nil {
		return err
	}

	provider = &PostgresProvider{DB: db, TableName: "comments"}

	err = cleanDB(db)
	if err != nil {
		return err
	}

	return nil
}

func TestGlobal(t *testing.T) {
	// connectDB()
	// defer closeDB(provider.DB)
	//
	// t.Run("New", testNew)
	// t.Run("List", testList)
	// t.Run("update", testUpdate)
	// t.Run("delete", testDelete)
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
	comments, err := provider.List(context.Background(), 1, commentNum, 0)

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
	err := provider.Update(context.Background(), 1, "updated message")

	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
}

// Test delete
func testDelete(t *testing.T) {
	err := provider.Delete(context.Background(), 1)
	if err != nil {
		t.Error("Error should be nil but got:", err)
		return
	}
}
