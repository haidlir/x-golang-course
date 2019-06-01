package main

import (
	"fmt"
)

type PersegiPanjang struct {
	Panjang int
	Lebar int
}

func (Bidang PersegiPanjang) Luas() int {
	return Bidang.Panjang * Bidang.Lebar
}

func (Bidang PersegiPanjang) Keliling() int {
	return (Bidang.Panjang + Bidang.Lebar) * 2
}

func main() {
	pp := PersegiPanjang{
		Panjang: 8,
		Lebar: 6,
	}
	fmt.Println("Luas Bidang: ", pp.Luas())
	fmt.Println("Keliling Bidang: ", pp.Keliling())
}