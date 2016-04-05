package main

import (
    "bufio"
    "fmt"
    "os"
    "time"
)

func get(out chan string) {
    reader := bufio.NewReader(os.Stdin)
    result, _, _ := reader.ReadLine()
    out <- string(result)
}

func timeout(out chan int) {
    time.Sleep(time.Duration(time.Duration.Seconds(10000)))
    out <- 1
}

func main() {
    echo := make(chan string)
    quit := make(chan int)
    go get(echo)
    go timeout(quit)
    select {
    case result := <-echo:
        fmt.Println(result)
    case <-quit:
        fmt.Println("GOTTA GO FAST BRO")
    }
}
