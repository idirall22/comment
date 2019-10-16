package comment

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	pr "github.com/idirall22/comment/providers/postgres"
	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "diskshar_test"
)

var testService *Service

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

// Connect to db test
func connectDB() error {

	dbInfos := fmt.Sprintf(`host=%s port=%d user=%s password=%s dbname=%s sslmode=disable`,
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", dbInfos)
	if err != nil {
		return err
	}

	provider := &pr.PostgresProvider{DB: db, TableName: "comments"}
	testService = &Service{
		provider: provider,
	}
	err = cleanDB(db)
	if err != nil {
		return err
	}

	testService = StartService(db)
	return nil
}

func TestGlobal(t *testing.T) {
	if err := connectDB(); err != nil {
		log.Fatal("Error connect database test, ", err)
	}

	defer closeDB(testService.provider.(*pr.PostgresProvider).DB)

	t.Run("add comment", testAddComment)
	t.Run("update a comment", testUpdateComment)
	t.Run("delete a comment", testDeleteComment)
}
