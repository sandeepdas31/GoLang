package main

import (
	"fmt"
)

func main() {
	sum:=0
	for i:=1; i<6; i++{
	if i%2==0{
	continue
	}
	sum+=i
	}
	fmt.Println(sum)
}
