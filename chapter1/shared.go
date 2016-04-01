package main

import (
    "fmt"
    "time"
    "math"
)

func normal_is_prime(primes []int, candidate int) bool {
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

func is_prime(primes []int, candidate *int, result chan bool, quit *bool) {
    square_root := math.Sqrt(float64(*candidate))
    for _, prime := range primes {
        if *quit {
            return
        }
        if square_root < float64(prime) {
            break
        }
        if *candidate%prime == 0 {
            *quit = true
            result <- false
        }
    }
    if !*quit {
        result <- true
    }
}

func worker(prime_channel <-chan []int, candidate *int, result chan bool, quit *bool) {
    for prime_slice := range prime_channel {
        fmt.Println(prime_slice)
        is_prime(prime_slice, candidate, result, quit)
    }
}

func primes(ceiling int) {
    primes := make([]int, ceiling)
    primes[0] = 2
    candidate := 3
    found := 1
    quit := false
    prime_channel := make(chan []int, 4)
    result := make(chan bool, 4)
    for i := 0; i < 4; i++ {
        go worker(prime_channel, &candidate, result, &quit)
    }
    for found < ceiling {
        if found < 4 {
            if normal_is_prime(primes, candidate) {
                primes[found] = candidate
                found++
            }
            candidate += 2
        } else {
            quit := false
            if quit {

            }
            result := make(chan bool, 4)
            bottom_chunk := 0
            top_chunk := found / 4
            for i := 0; i < 4; i++ {
                prime_channel <- primes[bottom_chunk:top_chunk]
                bottom_chunk = top_chunk
                top_chunk = top_chunk + found/4
            }
            is_prime := true
            for r := range result {
                if !r {
                    is_prime = false
                    close(result)
                }
            }
            if is_prime {
                primes[found] = candidate
                found++
            }
            candidate += 2
        }
    }
    close(result)
    close(prime_channel)
    fmt.Println(primes[999])
}


func main() {
    start := time.Now()
    primes(1000)
    fmt.Println(time.Now().Sub(start))
}