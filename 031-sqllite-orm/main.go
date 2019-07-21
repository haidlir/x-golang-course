package main

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite" // sqlite for gorm
)

// Siswa holds each siswa atributes
type Siswa struct {
	gorm.Model
	Nama  string `json:"nama"`
	Kelas int    `json:"kelas"`
}

func main() {
	ormDB, err := gorm.Open("sqlite3", "test.db")
	if err != nil {
		log.Printf("Unable to open sqllite DB: %v", err)
		return
	}
	// Migrate Schema
	ormDB.AutoMigrate(&Siswa{})

	log.Println("Add New Siswa")
	{
		newSiswa := new(Siswa)
		newSiswa.Nama = "Samuel"
		newSiswa.Kelas = 4
		ormDB.Create(newSiswa)
	}
	{
		newSiswa := new(Siswa)
		newSiswa.Nama = "Tetew"
		ormDB.Create(newSiswa)
	}

	log.Println("Print All Siswa")
	{
		allSiswa := []Siswa{}
		ormDB.Find(&allSiswa)
		fmt.Println(allSiswa)
	}
}
