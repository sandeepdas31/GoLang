package main

import (
	"fmt"
)

func main() {
	sum:=1
	
	switch sum{
	case 1:
		fmt.Println("one")
		fallthrough
	case 3:
		fmt.Println("Three")
	case 5:
		fmt.Println("Five")
		}
	
}
