package main

import "fmt"

const (
	A = iota
	KB = 11 << (iota*10)
	MB = 1 << (iota*10)
	GB = 1 << (iota*10)
)


func main(){
	fmt.Println("Decimal\t\tBinary")
	fmt.Printf("%d \t\t", KB)
	fmt.Printf("%b \n", KB)
	fmt.Printf("%d \t", MB)
	fmt.Printf("%b \n", MB)
	fmt.Printf("%d \t", GB)
	fmt.Printf("%b \n", GB)

}
