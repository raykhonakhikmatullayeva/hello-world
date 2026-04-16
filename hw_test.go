package main

import (
	"fmt"
)

func Example_mask() {
	fmt.Println(Mask("1233435436312"))
	fmt.Println(Mask("8888888888888888"))
	fmt.Println(Mask("abdsderfvderfswe"))
	//Output:
	//1233 **** **** 2
	//
	//8888 **** **** 8888
	//abds **** **** fswe
}
