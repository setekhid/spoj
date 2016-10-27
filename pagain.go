package main

import (
	"fmt"
	"math"
	"time"
)

func main() {

	begin := time.Now().UnixNano()
	for i := 0; i < 20; i++ {
		max := int64(0x1<<32) - 1
		k := math.Sqrt(float64(max))
		base := Primes(nil, 2, int64(k)+1)

		for num := max - 1; num >= 2; num-- {
			if IsPrime(base, num) {
				fmt.Println(num)
				break
			}
		}
	}
	end := time.Now().UnixNano()
	println((end - begin) / 20 / 1000000)

	var T int
	fmt.Scanln(&T)
	for ; T > 0; T-- {

		var N int64
		fmt.Scanln(&N)

		k := math.Sqrt(float64(N))
		base := Primes(nil, 2, int64(k)+1)

		for num := N - 1; num >= 2; num-- {
			if IsPrime(base, num) {
				fmt.Println(num)
				break
			}
		}
	}
}

func Primes(primes []int64, lower, upper int64) []int64 {

	for num := lower; num < upper; num++ {
		if IsPrime(primes, num) {
			primes = append(primes, num)
		}
	}
	return primes
}

func IsPrime(primes []int64, num int64) bool {

	for i := 0; i < len(primes); i++ {
		if num%primes[i] == 0 {
			return false
		}
	}
	return true
}
