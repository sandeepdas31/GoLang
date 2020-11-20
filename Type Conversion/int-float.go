package main

import (
	"fmt"
)

func main() {
	var x int = 23
	fmt.Printf("%v, %T\n",x,x)
	i := float32(x)
	fmt.Printf("%v, %T\n",i,i)
}
