package main

import "fmt"

type people interface{
	getName() 
	getAge() 
}

type emp struct{
	name string
	age int32
}

type stud struct{
	name string
	age int32
	rollNo int32
}

func main(){
	emp1 := emp{name: "arun", age: 24}
	stud1 := stud{name: "ram", age: 18, rollNo: 12}
	getDetails(emp1)
	getDetails(stud1)
}

func getDetails(v people){
	fmt.Println("values :", v)
	v.getName()
	v.getAge()
}

func (r emp) getName() {
	fmt.Println("the name of emp is", r.name)
}

func (r stud) getName() {
	fmt.Println("the name of emp is", r.name)
}

func (r emp) getAge() {
	fmt.Println("the name of emp is", r.age)
}

func (r stud) getAge() {
	fmt.Println("the name of emp is", r.age)
}