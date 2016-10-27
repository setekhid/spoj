package main

import (
	"fmt"
	"math/big"
)

func main() {

	primes := Primes(10000)

	var N int
	fmt.Scanln(&N)

	for N > 0 {

		sum := big.NewInt(1)

		for i := 0; i < len(primes) && primes[i] <= N; i++ {
			sum.Mul(sum, big.NewInt(2*CountFactors(N, primes[i])+1))
		}

		fmt.Println(sum.String())

		fmt.Scanln(&N)
	}
}

func Primes(n int) []int {

	primes := make([]int, 0)

	for num := 2; num <= n; num++ {

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

func CountFactors(factorial, factor int) int64 {

	sum := int64(0)
	for factorial >= factor {
		sum += int64(factorial / factor)
		factorial = factorial / factor
	}
	return sum
}
