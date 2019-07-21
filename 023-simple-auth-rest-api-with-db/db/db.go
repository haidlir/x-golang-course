package db

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite for gorm

	nmodel "github.com/haidlir/x-golang-course/023-simple-auth-rest-api-with-db/model"
)

// NewSQLLiteORMDB return the SQLLiteORM DB
func NewSQLLiteORMDB() *SQLLiteORM {
	ormDB, err := gorm.Open("sqlite3", "sqllite.db")
	if err != nil {
		log.Printf("Unable to open sqllite DB: %v", err)
		return nil
	}
	// Migrate Schema
	ormDB.AutoMigrate(&nmodel.Siswa{})
	db := new(SQLLiteORM)
	db.db = ormDB
	if ok := db.db.HasTable(&nmodel.Siswa{}); !ok {
		return nil
	}
	return db
}
