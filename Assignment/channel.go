package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
		
	go func() {
		fmt.Print("Hello")
		
		time.Sleep(time.Second * 2)

		
		<- done
	}()

	done <- struct{}{}
	fmt.Println(" world!")
}
