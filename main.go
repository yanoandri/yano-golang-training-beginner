package main

import (
	"fmt"
	"os"
	"reflect"
)

func main() {
	// TODO: task #1 - Why is it not working?
	// changing the type from const into var
	// the reason to change it because the value depends on env variables in a system
	// so using var will be the best option
	var task1 = os.Getenv("task_1")
	if !reflect.DeepEqual("", task1) {
		fmt.Println("Task #1 failed!")
	}
	// TODO: Challenge #1 - How to convert from any type into string
	// use FsprintF method to convert into string the simple way
	var challenge1 interface{} // DO NOT CHANGE THE DATATYPE
	challenge1 = 1
	challenge1 = fmt.Sprintf("%v", challenge1)
	if !reflect.DeepEqual("1", challenge1) {
		fmt.Println("Challenge #1 failed!")
	}
	fmt.Println("All passed!")
}
