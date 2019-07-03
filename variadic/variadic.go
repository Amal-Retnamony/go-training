package main

import "fmt"

func main(){
	sum(1,2)
	sum(1,2,3)
}

func sum(nums ...int){
	fmt.Println("Argument received",nums)
	for i,num := range(nums){
        fmt.Println("Index",i,"have value",num)
	}
} 