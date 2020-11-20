package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x int = 23
	fmt.Printf("%v, %T\n",x,x)
	i := strconv.Itoa(x)
	fmt.Printf("%v, %T\n",i,i)
}
