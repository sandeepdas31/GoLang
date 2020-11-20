package main

import (
	"fmt"
)

func main() {
	var x float32 = 23.4
	fmt.Printf("%v, %T\n",x,x)
	i := int(x)
	fmt.Printf("%v, %T\n",i,i)
}
