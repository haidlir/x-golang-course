package main

import (
	"fmt"
)

func main() {
	// iterasi 7 kali
	for i:=0; i<7; i++ {
		fmt.Println("i=", i)
	}
	fmt.Println("----------------")

	// iterasi slice
	himpunan := []int{1, 2, 3, 4, 5, 6, 7, 8, 19, 100}
	for i, val := range himpunan {
		fmt.Println("ke-", i, "adalah", val)
	}
	fmt.Println("----------------")

	// iterasi slice
	indexHalaman := map[string]int{
		"Judul": 1,
		"Bab 1": 10,
		"Bab 2": 25,
		"Bab 3": 50,
	}
	for key, val := range indexHalaman {
		fmt.Println("halaman", key, "adalah", val)
	}
	fmt.Println("----------------")
}