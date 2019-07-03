package main

import "fmt"


func main(){
	fmt.Println("Sum from 1 - 5 is", sum(5))
}


func sum(n int) int{
	if n==1 { 
		return 1
	}
	return n+sum(n-1)
}