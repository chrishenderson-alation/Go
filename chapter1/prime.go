package main

import (
	"fmt"
	"math"
	"time"
)

func is_prime(primes []int, candidate int) bool {
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

func is_prime_routine(result chan bool, primes []int, candidate int) {
	result <- is_prime(primes, candidate)
}

func primes(ceiling int) {
	primes := make([]int, ceiling)
	primes[0] = 2
	candidate := 3
	found := 1
	for found < ceiling {
		if found < ceiling {
			if is_prime(primes, candidate) {
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
			bottom_chunk := 0
			top_chunk := found / 4
			for i := 0; i < 4; i++ {
				go is_prime_routine(result, primes[bottom_chunk:top_chunk], candidate)
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
	fmt.Println(primes[ceiling-1])
}

func main() {

	start := time.Now()
	primes(300000)
	fmt.Println(time.Now().Sub(start))
}
