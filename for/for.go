package main

import "fmt"

func main(){
	for i:=1 ; ; i++ {
		if(i==2){
		  continue
		}
		fmt.Println(i)
		if(i == 5){
			break
		}
	}
    var j int = 0
	for j < 4{
		fmt.Println(j)
		j = j+1
	}
}