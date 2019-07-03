package main

import "fmt"

func main(){
   newInt := intSeq()
   fmt.Println(newInt())
   fmt.Println(newInt())
   fmt.Println(newInt())
}


func intSeq() func() int{
	i:= 0
	return func() int{
		i++
		return i
	}
}