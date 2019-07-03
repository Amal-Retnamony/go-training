package main

import "fmt"

type empType struct {
   name string
   age int
   role string
}


func main(){
	var emp1 empType
	emp1.name = "arun"
	emp1.age = 25
	emp1.role = "developer"
	fmt.Println("the details of employee 1", emp1)
	emp2 := empType{name: "ram", age: 24, role: "architect"}
	fmt.Println("the details of employee 2", emp2)
	emp3 := emp2
	fmt.Println("the details of employee 3", emp3)
}