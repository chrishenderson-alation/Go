package main

import "fmt"

func init() {
	fmt.Println("Me first!")
}

func init() {
	fmt.Println("Me second!")
}

func main() {
	fmt.Println("Hold on, who are those other guys? I'm the main guy here!")
}

func init() {
	fmt.Println("Me third!")
}
