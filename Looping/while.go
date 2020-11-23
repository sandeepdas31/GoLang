package main

import (
	"fmt"
)

func main() {
	sum:=0 
	i:=1
	for i<=5{
		sum+=i
		i++
	}
	fmt.Println(sum)
}
