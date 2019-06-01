package model

// Siswa holds each siswa atributes
type Siswa struct {
	ID    int    `json:"id"`
	Nama  string `json:"nama"`
	Kelas int    `json:"kelas"`
}
