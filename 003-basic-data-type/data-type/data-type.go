package main

import (
	"fmt"
)

func main() {
	// String
	text := "This is a string"
	fmt.Println(text)
	var textAlso string = "This is also a text"
	fmt.Println(textAlso)
	
	// Boolean
	isHuman := true
	isPythonCourse := false
	fmt.Println("Human ? ", isHuman)
	fmt.Println("Python Course ? ", isPythonCourse)

	// Numeric Integer
	a := 9999
	fmt.Printf("An Integer with default type %T: %v\n", a, a)
	// int8
	var b1 int8 = 127
	fmt.Println("An 8 bytes integer: ", b1)
	var b2 int8 = -128
	fmt.Println("An 8 bytes integer: ", b2)
	// uint8
	var c1 uint8 = 255
	fmt.Println("An 8 bytes unsigned integer: ", c1)
	

	// Numeric Decimal or Floating Point
	d1 := 256.0
	fmt.Printf("A Float with default type %T: %v\n", d1, d1)
	// Float32
	var d2 float32 = 300.
	fmt.Printf("A Float with default type %T: %v\n", d2, d2)
	// Float64
	var d3 float64 = 1000.
	fmt.Printf("A Float with default type %T: %v\n", d3, d3)

	// Complex Number
	e1 := 1.6 + 1.2i
	fmt.Printf("A complex number with default type %T: %v\n", e1, e1)
	// complex128
	var e2 complex64 = 4 + 3i
	fmt.Printf("A complex number with default type %T: %v\n", e2, e2)
	// complex256
	var e3 complex128 = 4 + 3i
	fmt.Printf("A complex number with default type %T: %v\n", e3, e3)

	// Byte
	var f1 byte = 255
	fmt.Printf("A byte with default type %T: %v\n", f1, f1)
	f1++
	fmt.Printf("A byte with default type %T: %v\n", f1, f1)

	// Rune
	var g1 rune = '�'
	fmt.Printf("A rune with default type %T: %U\n", g1, g1)
	fmt.Println(rune(g1))
}