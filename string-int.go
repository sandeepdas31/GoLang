package main

import (
	"fmt"
)

func main() {
	var x string = "hi"
	fmt.Printf("%v, %T\n",x,x)
	i := []byte(x)
	fmt.Printf("%v, %T\n",i,i)
}
