package main

import "fmt"

func main(){
	var s = make([]int, 3)
	s[0] = 1
	s[1] = 2
	s[2] = 3
	fmt.Println("the slice s is", s)
	fmt.Println("the length of s is", len(s))
	s = append(s, 4)
	fmt.Println("Now the slice s is", s)
	fmt.Println("Now the length of s is", len(s))
}