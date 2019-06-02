package db

import (
	"fmt"
	"sync"

	nmodel "github.com/haidlir/x-golang-course/021-simple-rest-api/model"
)

// DummyDB is an inMem memory
type DummyDB struct {
	daftarSiswa []nmodel.Siswa
	idSiswa     int
	mux         sync.Mutex
}

// GetAllSiswa returns all siswa
func (db *DummyDB) GetAllSiswa() []nmodel.Siswa {
	db.mux.Lock()
	defer db.mux.Unlock()
	return db.daftarSiswa
}

// GetDetailSiswa returns specific siswa
func (db *DummyDB) GetDetailSiswa(id int) *nmodel.Siswa {
	db.mux.Lock()
	defer db.mux.Unlock()
	index, err := db.findID(id)
	if err != nil {
		return nil
	}
	return &db.daftarSiswa[index]
}

// AddSiswa adds new siswa to the DB and returns the error status
func (db *DummyDB) AddSiswa(siswaBaru nmodel.Siswa) (*nmodel.Siswa, error) {
	db.mux.Lock()
	siswaBaru.ID = db.idSiswa
	db.idSiswa++
	db.daftarSiswa = append(db.daftarSiswa, siswaBaru)
	db.mux.Unlock()
	return &siswaBaru, nil
}

// UpdateSiswa updates specific siswa and return the error status
func (db *DummyDB) UpdateSiswa(id int, data nmodel.Siswa) (*nmodel.Siswa, error) {
	db.mux.Lock()
	defer db.mux.Unlock()
	index, err := db.findID(id)
	if err != nil {
		return nil, fmt.Errorf("id not found")
	}
	data.ID = id
	db.daftarSiswa[index] = data
	return &db.daftarSiswa[index], nil
}

// DeleteSiswa deletes specific siswa and returns error status
func (db *DummyDB) DeleteSiswa(id int) error {
	db.mux.Lock()
	defer db.mux.Unlock()
	index, err := db.findID(id)
	if err != nil {
		return fmt.Errorf("id not found")
	}
	db.daftarSiswa = append(db.daftarSiswa[:index], db.daftarSiswa[index+1:]...)
	return nil
}

func (db *DummyDB) findID(id int) (index int, err error) {
	for i, val := range db.daftarSiswa {
		if val.ID == id {
			index, err = i, nil
			return
		}
	}
	index, err = -1, fmt.Errorf("id not found")
	return
}
