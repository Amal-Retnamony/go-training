package main

import "fmt"

func main(){
	var s [5]int
	s[1] = 1
	s[2] = 2
	fmt.Println("s[1]", s[1])
	fmt.Println("length of array", len(s))
	j := [5]int{1,2,3,4,5}
	fmt.Println("Array j is", j)
}