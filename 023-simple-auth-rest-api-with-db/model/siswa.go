package model

import (
	"github.com/jinzhu/gorm"
)

// Siswa holds each siswa atributes
type Siswa struct {
	gorm.Model
	Nama  string `json:"nama"`
	Kelas int    `json:"kelas"`
}
