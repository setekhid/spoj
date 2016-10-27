package main

import (
	"fmt"
)

func main() {

	var N int
	fmt.Scanln(&N)

	nums := make([]int32, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}

	var Q int32
	fmt.Scanln(&Q)

	for ; Q > 0; Q-- {

		var op string
		fmt.Scan(&op)

		if op == "U" {

			var x int
			var y int32
			fmt.Scanln(&x, &y)
		} else {

			var x, y int
			fmt.Scanln(&x, &y)
		}
	}
}
