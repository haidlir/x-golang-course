package main

import (
	"fmt"
)

func main() {
	// array
	himpunan := [10]int{11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	fmt.Println(himpunan)
	himpunan[3] = 24
	fmt.Println("himpunan ke 4 adalah", himpunan[3])
	fmt.Println(himpunan)

	// slice
	kumpulanNama := []string{"andi", "budi", "candra"}
	fmt.Println(kumpulanNama)
	// tambah
	kumpulanNama = append(kumpulanNama, "deni")
	fmt.Println(kumpulanNama)
	// hapus
	kumpulanNama = kumpulanNama[:len(kumpulanNama)-1]
	fmt.Println(kumpulanNama)

	// map
	indexHalaman := map[string]int{
		"Judul": 1,
		"Bab 1": 10,
		"Bab 2": 25,
		"Bab 3": 50,
	}
	fmt.Println(indexHalaman)
	fmt.Println(indexHalaman["Bab 2"])
	// tambah
	indexHalaman["Bab 5"] = 100
	fmt.Println(indexHalaman)
	// hapus
	delete(indexHalaman, "Bab 3")
	fmt.Println(indexHalaman)
}