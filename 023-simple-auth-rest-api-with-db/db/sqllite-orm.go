package db

import (
	"fmt"

	"github.com/jinzhu/gorm"

	nmodel "github.com/haidlir/x-golang-course/023-simple-auth-rest-api-with-db/model"
)

// SQLLiteORM is a DB connector using ORMs
type SQLLiteORM struct {
	db *gorm.DB
}

// Close the DB Connection
func (db *SQLLiteORM) Close() {
	db.db.Close()
}

// GetAllSiswa returns all siswa
func (db *SQLLiteORM) GetAllSiswa() []nmodel.Siswa {
	var siswas []nmodel.Siswa
	db.db.Find(&siswas)
	return siswas
}

// GetDetailSiswa returns specific siswa
func (db *SQLLiteORM) GetDetailSiswa(id int) *nmodel.Siswa {
	var siswa nmodel.Siswa
	db.db.First(&siswa, id)
	if siswa.ID != uint(id) {
		return nil
	}
	return &siswa
}

// AddSiswa adds new siswa to the DB and returns the error status
func (db *SQLLiteORM) AddSiswa(siswaBaru nmodel.Siswa) (*nmodel.Siswa, error) {
	db.db.Create(&siswaBaru)
	return &siswaBaru, nil
}

// UpdateSiswa updates specific siswa and return the error status
func (db *SQLLiteORM) UpdateSiswa(id int, data nmodel.Siswa) (*nmodel.Siswa, error) {
	var siswa nmodel.Siswa
	db.db.First(&siswa, id)
	siswa.Nama = data.Nama
	siswa.Kelas = data.Kelas
	db.db.Save(&siswa)
	return &siswa, nil
}

// DeleteSiswa deletes specific siswa and returns error status
func (db *SQLLiteORM) DeleteSiswa(id int) error {
	if id <= 0 {
		return fmt.Errorf("No id %v", id)
	}
	var siswa nmodel.Siswa
	db.db.First(&siswa, id)
	if siswa.ID != uint(id) {
		return fmt.Errorf("unable to find id %v", id)
	}
	db.db.Delete(&siswa)
	return nil
}
