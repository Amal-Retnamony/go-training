package main

import "fmt"

func main(){
	 add := sum(1,2)
	 fmt.Println("1 + 2 = ",add)
	 result := mul(5, 10)
	 fmt.Println("5 * 10 = ",result)
}

func sum( a int,  b int) int {
	return a+b
}

func mul( a,b int) int {
	return a*b
}