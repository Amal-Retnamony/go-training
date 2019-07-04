package main

import "fmt"

type emp struct{
	name string
	age int
	role string
}




func main(){
	emp1 := emp{name: "arun", age: 25, role: "developer"}
	fmt.Println("the name of employee 1 is", emp1.getName())
	fmt.Println("the age of employee 1 is", emp1.getAge())
}

func (e emp)getName() string{
	return e.name
} 

func (e *emp)getAge() int{
	return e.age
} 