package db_test

import (
	"testing"

	ndb "github.com/haidlir/x-golang-course/021-simple-rest-api/db"
	nmodel "github.com/haidlir/x-golang-course/021-simple-rest-api/model"
)

func TestSiswa(t *testing.T) {
	db := ndb.NewDummyDB()
	// Add Siswa
	t.Run("Add Siswa", func(t *testing.T) {
		siswaBaruDummy := nmodel.Siswa{
			Nama:  "dummy",
			Kelas: 0,
		}
		siswaBaru := nmodel.Siswa{
			Nama:  "Ardi",
			Kelas: 6,
		}
		_, _ = db.AddSiswa(siswaBaruDummy)
		result, err := db.AddSiswa(siswaBaru)
		if err != nil {
			t.Fatalf("Error while adding siswa baru: %v", err)
		}
		if result == nil {
			t.Fatalf("result is nil")
		}
		if result.ID != 1 {
			t.Errorf("id result should be 1 instead of %v", result.ID)
		}
		if result.Nama != siswaBaru.Nama {
			t.Errorf("nama in result should be %v instead of %v", siswaBaru.Nama, result.Nama)
		}
		if result.Kelas != siswaBaru.Kelas {
			t.Errorf("kelas in result should be %v instead of %v", siswaBaru.Kelas, result.Kelas)
		}
	})
	// Get All
	t.Run("Get All Siswa", func(t *testing.T) {
		daftarSiswa := db.GetAllSiswa()
		if len(daftarSiswa) != 2 {
			t.Fatalf("The amount of siswa should be 2 instead of %v", len(daftarSiswa))
		}
		if daftarSiswa[1].Nama != "Ardi" {
			t.Fatalf("Nama in daftarSiswa[1] should be Ardi instead of %v", daftarSiswa[1].Nama)
		}
		if daftarSiswa[1].Kelas != 6 {
			t.Fatalf("Kelas in daftarSiswa[1] should be Ardi instead of %v", daftarSiswa[1].Kelas)
		}
	})
	// Get Specific Siswa
	t.Run("Get Specific Siswa", func(t *testing.T) {
		id := 1
		siswa := db.GetDetailSiswa(id)
		if siswa == nil {
			t.Fatal("siswa should be not nil")
		}
		if siswa.Nama != "Ardi" {
			t.Fatalf("Nama in siswa should be Ardi instead of %v", siswa.Nama)
		}
		if siswa.Kelas != 6 {
			t.Fatalf("Kelas in siswa should be Ardi instead of %v", siswa.Kelas)
		}
	})
	// Update Specific Siswa
	t.Run("Update Specific Siswa", func(t *testing.T) {
		id := 1
		siswa := db.GetDetailSiswa(id)
		siswa.Nama = "Arman"
		siswa.Kelas = 7
		updatedSiswa, err := db.UpdateSiswa(id, *siswa)
		if err != nil {
			t.Fatalf("Updating siswa returns error: %v", err)
		}
		if updatedSiswa == nil {
			t.Fatal("updatedSiswa should be not nil")
		}
		if updatedSiswa.Nama != siswa.Nama {
			t.Fatalf("Nama in updatedSiswa should be %v instead of %v", siswa.Nama, updatedSiswa.Nama)
		}
		if updatedSiswa.Kelas != siswa.Kelas {
			t.Fatalf("Kelas in updatedSiswa should be %v instead of %v", siswa.Kelas, updatedSiswa.Kelas)
		}
		updatedSiswa = db.GetDetailSiswa(id)
		if updatedSiswa == nil {
			t.Fatal("updatedSiswa should be not nil")
		}
		if updatedSiswa.Nama != siswa.Nama {
			t.Fatalf("Nama in updatedSiswa should be %v instead of %v", siswa.Nama, updatedSiswa.Nama)
		}
		if updatedSiswa.Kelas != siswa.Kelas {
			t.Fatalf("Kelas in updatedSiswa should be %v instead of %v", siswa.Kelas, updatedSiswa.Kelas)
		}
	})
	// Delete Specific Siswa
	t.Run("Delete Specific Siswa", func(t *testing.T) {
		id := 0
		err := db.DeleteSiswa(id)
		if err != nil {
			t.Fatalf("delete specific siswa error: %v", err)
		}
		daftarSiswa := db.GetAllSiswa()
		if len(daftarSiswa) != 1 {
			t.Fatalf("The amount of siswa should be 1 instead of %v", len(daftarSiswa))
		}
		deleletedSiswa := db.GetDetailSiswa(id)
		if deleletedSiswa != nil {
			t.Fatalf("Siswa with id %v should be not found", id)
		}
	})
}
