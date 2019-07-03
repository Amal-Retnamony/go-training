package main

import "fmt"

func main(){
   i := 0
   zeroValue(i)
   fmt.Println("zeroValue", i) 	
   zeroPointer(&i)
   fmt.Println("zeroPointer",i) 
}

func zeroValue(a int) {
	a= 1
}

func zeroPointer(a *int) {
	*a = 1
}