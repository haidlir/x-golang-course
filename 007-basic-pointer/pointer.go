package main

import (
	"fmt"
)

func main() {
	var1 := 10
	fmt.Printf("var1: %v\n", var1)
	fmt.Printf("pointer dari var1: %v\n", &var1)
	fmt.Printf("tipe data pointer var1 %T\n", &var1)
}
