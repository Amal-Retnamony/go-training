package main


import "errors"
import "fmt"

type err struct{
    value int
	arg error
}

func main(){
	for _, v := range []int{1,2,3}{
		i,e := f1(v)
		if e!=nil{
			fmt.Println("error occured", e)
		}else{
			fmt.Println("successfull", i)
		}
	}
	_,e1 := f2(3)
	fmt.Println("error method",e1)
}

func f1(a int) (int,error) {
	if(a==3){
       return -1, errors.New("cannot process 3")
	}
	return a*10,nil
}

func f2(a int) (int,error) {
	if(a==3){
       return -1, &err{3,errors.New("cannot process 3")}
	}
	return a*10,nil
}

func (e *err) Error() string {
    return fmt.Sprintf("%d - %s", e.value, e.arg)
}

