package main

import (
	"fmt"
)

func main() {
	str:=[]string{"hello","people"}
	for i, s:= range str{
	fmt.Println(i,s)
	}
}
