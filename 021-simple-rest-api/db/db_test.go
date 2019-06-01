package db_test

import (
	"testing"

	ndb "github.com/haidlir/x-golang-course/021-simple-rest-api/db"
)

func TestNewDummy(t *testing.T) {
	db := ndb.NewDummyDB()
	if db == nil {
		t.Fatal("db should be not nil")
	}
}
