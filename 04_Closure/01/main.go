package main

import "fmt"

func wrappeer() func() int{
	x := 9
	return func() int{
		x++
		return x
	}
}


func main(){
	increment := wrappeer()
	fmt.Printf("%v \n", wrappeer())
	fmt.Printf("%T \n", wrappeer())

	fmt.Println(increment())
	againIncrement := wrappeer()
	fmt.Println(x)
	for i:=0; i<4; i++{
		againIncrement()
	}
	fmt.Println(againIncrement())
}
