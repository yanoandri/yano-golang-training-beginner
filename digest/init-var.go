package main

import (
	"errors"
	"fmt"
)

var variable1, variable2, variable3, variable4 = "var1", "var2", 3, true

const Pi = 3.14
const COUNTRY = "INDONESIA"
const PREFIX = "wp_"

func main() {
	// assign to shorthand only inside function
	shorthand1, shorthand2 := variable1, variable2
	// print the variable
	println(variable1)
	println(variable2)
	// using shorthand
	println(shorthand1)
	println(shorthand2)
	// print constants
	println(COUNTRY)
	println(PREFIX + "wordpress")
	// numbers
	number1, number2, number3 := 3, 0, 0.0
	var someFloat1 float64 = 5.6
	var someFloat2 float64 = 7.0
	number3 += someFloat1 + someFloat2
	number2 += number1 + variable3
	fmt.Printf("Result 1 is %v", number2)
	println()
	fmt.Printf("Result 2 is %v", number3)
	println()
	// combine 2 strings
	helloJpn1, helloJpn2, lower := "Hello", "Konichiwa", ""
	fmt.Println(helloJpn1 + " " + helloJpn2)
	// change the first to lower
	someBytes := []byte(helloJpn1)
	someBytes[0] = 'h'
	lower = string(someBytes)
	println(lower)
	// get index
	cut1 := "c" + helloJpn1[:1] + helloJpn1[4:5] // substring
	println(cut1)
	multiline := `test
	test2
	test3`
	println(multiline)
	// errors
	err := errors.New("emit macho dwarf: elf header corrupted")
	fmt.Println(err) // cannot print with println
	// some grouping variable
	// var (
	// 	grouping1 int
	// 	grouping2 string
	// 	grouping3 bool
	// )
	// iota
	const (
		iota1 = iota
		iota2 = iota
		iota3
	)
	println(iota1)
	println(iota2)
	println(iota3)
}
