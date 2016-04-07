package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func cat() {
	var s string
	var sep string
	for _, argument := range os.Args[1:] {
		s += sep + argument
		sep = " "
	}
	fmt.Println(s)
}

func join() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}

func main() {
	start := time.Now()
	cat()
	fmt.Println(time.Now().Sub(start))
	start = time.Now()
	join()
	fmt.Println(time.Now().Sub(start))

	// Yay, join!
	// 31.739µs -> string concatenation
	// 1.164µs -> strings.Join
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	fmt.Println(input.Text())
}
