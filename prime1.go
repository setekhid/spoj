package main

import (
	"fmt"
	"math"
)

func main() {

	var t int
	fmt.Scanln(&t)

	for ; t > 0; t-- {

		var m, n int32
		fmt.Scanln(&m, &n)

		k := math.Min(math.Sqrt(float64(n)), float64(m-1))
		primes := Primes(nil, 2, int32(k)+1)
		mi := len(primes)
		if m < 2 {
			primes = Primes(primes, 2, n+1)
		} else {
			primes = Primes(primes, m, n+1)
		}

		for i := mi; i < len(primes); i++ {
			fmt.Println(primes[i])
		}

		fmt.Println()
	}
}

func Primes(primes []int32, lower, upper int32) []int32 {

	for num := lower; num < upper; num++ {

		isPrime := true
		for i := 0; i < len(primes); i++ {
			if num%primes[i] == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			primes = append(primes, num)
		}
	}

	return primes
}
