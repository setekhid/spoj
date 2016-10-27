package main

import (
	"fmt"
	"math"
)

func main() {

	var S, N int
	fmt.Scanln(&S, &N)

	ss := make([]int, N)
	vs := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&ss[i], &vs[i])
	}

	stats := make([]int, S+1)
	for i := 0; i < N; i++ {
		for Si := 0; Si <= S; Si++ {

			V := stats[Si]
			if Si >= ss[i] {
				stats[Si-ss[i]] = int(math.Max(float64(stats[Si-ss[i]]), float64(V+vs[i])))
			}
		}
	}

	max := 0
	for _, V := range stats {
		if V > max {
			max = V
		}
	}

	fmt.Println(max)
}
