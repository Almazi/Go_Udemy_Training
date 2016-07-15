package main

import "fmt"

func main(){
	for i := 0; i < 30; i++{
		fmt.Printf("%d \t %b \t %x \n", i , i, i)
	}
}
