package main

import (
    "fmt"
    "reflect"
)

func main() {
    var numerator float64
    numerator = 5.0
    divisor := 2.0
    fmt.Println("Hello, Go! Happy to be here!")
    fmt.Println("The type of", numerator, "is", reflect.TypeOf(numerator))
    fmt.Println("The type of", divisor, "is", reflect.TypeOf(divisor))
    fmt.Println(numerator, "/", divisor, "is", numerator/divisor)
}
