package comment

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	pr "github.com/idirall22/comment/providers/postgres"
	u "github.com/idirall22/user"
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

var testToken string
var userUsernameTest = "alice"
var userPasswordTest = "fdpjfd654/*sMLdf"

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

// Get user auth allows to login and get a test token
func getUserAuth(db *sql.DB) {

	m := make(map[string]string)
	m["username"] = userUsernameTest
	m["password"] = userPasswordTest

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println(err)
		return
	}
	body := bytes.NewReader(b)

	serviceUser := u.StartService(db, "users")

	w := httptest.NewRecorder()
	r, err := http.NewRequest("POST", "/login", body)

	if err != nil {
		fmt.Println(err)
		return
	}

	h := http.HandlerFunc(serviceUser.Login)
	h.ServeHTTP(w, r)

	testToken = w.Header().Get("Autherization")
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

	getUserAuth(db)

	testService = StartService(db)
	return nil
}

func TestGlobal(t *testing.T) {
	if err := connectDB(); err != nil {
		log.Fatal("Error connect database test, ", err)
	}

	defer closeDB(testService.provider.(*pr.PostgresProvider).DB)

	// t.Run("add comment", testAddComment)
	// t.Run("update a comment", testUpdateComment)
	// t.Run("delete a comment", testDeleteComment)

	t.Run("add a comment handler", testAddCommentHandler)
	t.Run("update a comment handler", testUpdateCommentHandler)
	t.Run("delete a comment handler", testDeleteCommentHandler)
}
