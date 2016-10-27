package main

import (
	"fmt"
)

/*
 *  this implementation has LTS. see the alternative in haskell dotaa.hs
 */
func main() {

	var t int
	fmt.Scanln(&t)

	for ; t > 0; t-- {

		var n, m, D, H int
		fmt.Scanln(&n, &m, &D)

		passed := 0
		heroi := 0

		for ; heroi < n && passed < m; heroi++ {
			fmt.Scanln(&H)
			passed += (H - 1) / D
		}

		for ; heroi < n; heroi++ {
			var dummy string
			fmt.Scanln(&dummy)
		}

		if passed >= m {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
