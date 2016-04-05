package main

import (
    "fmt"
    "math"
    "runtime"
    "time"
)

type Work struct {
    candidate int
    primes    []int
    is_prime  chan bool
}

func primality_test(work Work) {
    square_root := math.Sqrt(float64(work.candidate))
    for _, prime := range work.primes {
        if float64(prime) > square_root {
            work.is_prime <- true
            return
        }
        if work.candidate%prime == 0 {
            work.is_prime <- false
            return
        }
    }
}

func test_primes(ceiling int) {
    primes := make([]int, ceiling)
    primes[0] = 2
    candidate := 3
    found := 1
    work := make([]Work, runtime.NumCPU())
    for i := 0; i < runtime.NumCPU(); i++ {
        work[i] = Work{candidate, primes, make(chan bool, 1)}
    }
    worker := work[0]
    for found < 100 {
        worker.candidate = candidate
        primality_test(worker)
        select {
        case result := <-worker.is_prime:
            if result {
                primes[found] = candidate
                worker.primes = primes
                found++
            }
        }
        candidate += 2
    }
    for {
        for i := 0; i < len(work); i++ {
            work[i].candidate = candidate
            work[i].primes = primes
            candidate += 2
            go primality_test(work[i])
        }
        for i := 0; i < len(work); i++ {
            select {
            case result := <-work[i].is_prime:
                if result {
                    primes[found] = work[i].candidate
                    found++
                    if found >= ceiling {
                        fmt.Println(primes[ceiling-1])
                        return
                    }
                }
            }
        }
    }
}

func main() {
    start := time.Now()
    test_primes(10000)
    fmt.Println(time.Now().Sub(start))
}
