package main

import (
    "fmt"
    "math"
    "time"
)

func test(primes []int, candidate int) bool {
    square_root := math.Sqrt(float64(candidate))
    for _, prime := range primes {
        if square_root < float64(prime) {
            break
        }
        if candidate%prime == 0 {
            return false
        }
    }
    return true
}

func test_routine(primes []int, candidate int, result chan bool, quit chan bool) func() {
    return func() {
        square_root := math.Sqrt(float64(candidate))
        for _, prime := range primes {
            select {
            case <-quit:
                quit <- true
                return
            default:
            }
            if square_root < float64(prime) {
                break
            }
            if candidate%prime == 0 {
                result <- false
                quit <- true
                return
            }
        }
        result <- true
    }
}

func worker(functions chan func()) {
    for function := range functions {
        function()
    }
}

func primes(ceiling int) {
    primes := make([]int, ceiling)
    primes[0] = 2
    candidate := 3
    found := 1
    functions := make(chan func(), 4)
    for i := 0; i < 4; i++ {
        go worker(functions)
    }
    for found < ceiling {
        if found < 100 {
            if test(primes, candidate) {
                primes[found] = candidate
                found++
            }
            candidate += 2
        } else {
            /*
               WELL, this divide and conquer works, but the overhead in spinning
               up goroutines is still too much.
            */
            result := make(chan bool, 4)
            quit := make(chan bool, 4)
            bottom_chunk := 0
            top_chunk := found / 4
            for i := 0; i < 4; i++ {
                prime_slice := primes[bottom_chunk:top_chunk]
                functions <- test_routine(prime_slice, candidate, result, quit)
                bottom_chunk = top_chunk
                top_chunk = top_chunk + found/4
            }
            is_prime := true
            for i := 0; i < 4; i++ {
                if !<-result {
                    is_prime = false
                    break
                }
            }
            if is_prime {
                primes[found] = candidate
                found++
            }
            candidate += 2
        }
    }
    fmt.Println(primes[999])
}

func main() {
    start := time.Now()
    primes(1000)
    fmt.Println(time.Now().Sub(start))
}
