package main

import (
	"fmt"
	"Go_Udemy_Training/02_Packages/stringUtl"
)

func main(){
	fmt.Println(stringUtl.Reverse("jeera pani")) //Accessing Reverse
	// function of stringUtl package. As Reverse starts with capital R it
	// is importable.
	fmt.Println(stringUtl.MyName)
}

