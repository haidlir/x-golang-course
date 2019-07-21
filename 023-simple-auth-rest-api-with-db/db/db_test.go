package db_test

import (
	"testing"

	ndb "github.com/haidlir/x-golang-course/023-simple-auth-rest-api-with-db/db"
)

func TestNewSQLLiteORMDB(t *testing.T) {
	db := ndb.NewSQLLiteORMDB()
	if db == nil {
		t.Fatal("db should be not nil")
	}
}
