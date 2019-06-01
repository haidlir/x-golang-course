package main

import "fmt"

func main() {
	sisiPersegi := 10
	luasPersegi := hitungLuasPersegi(sisiPersegi)
	fmt.Println("Luas Persegi adalah", luasPersegi, "cm2")
	kelilingPersegi := hitungKelilingPersegi(sisiPersegi)
	fmt.Println("Luas Persegi adalah", kelilingPersegi, "cm")
}

func hitungLuasPersegi(s int) int {
	luas := s * s
	return luas
}

func hitungKelilingPersegi(s int) int {
	keliling := s * 4
	return keliling
}
