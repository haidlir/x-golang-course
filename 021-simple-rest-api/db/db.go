package db

// NewDummyDB return the dummy DB
func NewDummyDB() *DummyDB {
	db := new(DummyDB) // &DummyDB{}
	return db
}
