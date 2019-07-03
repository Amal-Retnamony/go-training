package main

import "fmt"

func main(){
	sum,mul := calc(2,3)
	fmt.Println("2 + 3 =",sum)
	fmt.Println("2 * 3 =",mul)
}

func calc(a,b int)(int, int){
	return 2+3,2*3
}