package main

import "fmt"

var a = b + 1
var b = c + 1
var c = getOne()

func main() {
	fmt.Println(a, b, c)
}

func getOne() int {
	return 1
}
