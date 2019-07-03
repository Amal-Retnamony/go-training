package main

import "fmt"


func main(){
	var m = make(map[string]string)
	m["name"] = "arun"
	m["city"] = "trivandrum"
	fmt.Println("the map is ", m)
	delete(m, "city")
	fmt.Println("the map is ", m)
	_, prs := m["city"]
	fmt.Println("is value for city present: ", prs)
}