package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Domicile struct {
	country   string `json:"country"`
	is_remote bool   `json:"is_remote"`
}

type Employee struct {
	name            string   `json:"name"`
	entity          string   `json:"entity"`
	employee_number int      `json:"employee_number"`
	salary          float64  `json:"salary"`
	domicile        Domicile `json:"domicile"`
}

func main() {
	// TODO: task #1 - give me a skeleton!
	//missing comma after salary
	data := string(`
    {
        "name": "Golang",
        "entity": "Xendit",
        "employee_number": 10,
        "salary": 1.5,
        "domicile": {
            "country": "ID",
            "is_remote": true
        }
    }
    `)
	var employee Employee
	if err := json.Unmarshal([]byte(data), &employee); err != nil {
		fmt.Println("Task #1 failed!")
		return
	}
	// TODO: task #2 - I am a legal employee, include me into your database!
	var database map[string]Employee
	database = make(map[string]Employee) //use make map into employee
	database["Golang"] = employee        // assign employee with "Golang" Key hash map
	if !reflect.DeepEqual(database["Golang"], employee) {
		fmt.Println("Task #2 failed!")
		return
	}
	return
}
