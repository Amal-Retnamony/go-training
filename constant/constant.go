package main

import "fmt"


const s string = "constant"


func main(){
	fmt.Println("constant s is a ",s)
	const a = 5
	fmt.Println("constant a is", a)
}