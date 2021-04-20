package main

import (
	"fmt"
)

func main() {
	// basic array
	var someArray [10]int
	someArray[0] = 1
	someArray[9] = 9
	fmt.Printf("First element is : %v", someArray[0])
	println()
	fmt.Printf("Last element is : %v", someArray[9])
	println()
	// defining an array always print 1 element
	define1 := [3]int{1, 2, 3} // define an int array with 3 elements
	println(define1[0])
	define2 := [10]int{1, 2, 3}
	// define a int array with 10 elements, of which the first three are assigned.
	//The rest of them use the default value 0.
	println(define2[0])
	define3 := [...]int{4, 5, 6} // use `…` to replace the length parameter and Go will calculate it
	println(define3[0])
	// define 2 dimensional array and print 1 on 1
	multidimensional1 := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{4, 5, 6, 7}}
	multidimensional2 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Printf("%v pair is %v", multidimensional1[0][1], multidimensional1[1][2])
	println()
	fmt.Printf("%v pair is %v", multidimensional2[0][1], multidimensional2[1][2])
	println()
	// array slice
	ingredients := [...]string{"tomato", "potato", "banana"}
	slicingArray := ingredients[1:3]
	println(slicingArray[0])
	println(slicingArray[1])
	// map array
	someMap := make(map[string]int)
	someMap["ONE"] = 1
	someMap["TWO"] = 2
	someMap["THREE"] = 3
	newMap := new([2]int)
	println(someMap["ONE"])
	println(someMap["TWO"])
	println(someMap["THREE"])
	println(newMap)
	println(newMap[0])
}
