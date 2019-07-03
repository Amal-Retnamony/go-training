package main

import "fmt"
import "time"

func main(){
	var i int32 = 2
     switch i{
	 case 1:
		 fmt.Println("one")
	 case 2:
		fmt.Println("two")
	 case 3:
		 fmt.Println("three")
	 default:
		 fmt.Println("default")
	 }
	switch time.Now().Weekday(){
	case time.Sunday, time.Saturday: 
		fmt.Println("its is weekend")
	default:
		fmt.Println("its is not weekend")
	}

}