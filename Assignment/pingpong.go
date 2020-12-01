package main

import (
    "fmt"
    "time"
    "os"
)


func pinger(pinger <-chan int, ponger chan<- int) {
    for {
        val := <-pinger
        fmt.Println("ping")
        time.Sleep(500 * time.Millisecond)
        ponger <- val + 1
    }
}


func ponger(pinger chan<- int, ponger <-chan int) {
    for {
        <-ponger
        fmt.Println("pong")
        time.Sleep(200 * time.Millisecond)
        pinger <- 1
    }
}

func main() {
    ping := make(chan int)
    pong := make(chan int)

    go pinger(ping, pong)
    go ponger(ping, pong)

   
    ping <- 1

    time.Sleep(3 * time.Second)
    os.Exit(0)
}