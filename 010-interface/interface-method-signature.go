package main

import (
	"fmt"
)

type Bidang interface {
	Luas() int
	Keliling() int
}

func PrintBidang(b Bidang) {
	fmt.Printf("Bidang %T\n", b)
	fmt.Println("Luas:", b.Luas())
	fmt.Println("Keliling:", b.Keliling())
}

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

type Persegi struct {
	Sisi int
}

func (Bidang Persegi) Luas() int {
	return Bidang.Sisi * Bidang.Sisi
}

func (Bidang Persegi) Keliling() int {
	return Bidang.Sisi * 4
}

func main() {
	pp := PersegiPanjang{8, 6}
	p := Persegi{4}
	PrintBidang(pp)
	PrintBidang(p)
}