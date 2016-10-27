package main

import (
	"fmt"
	"sort"
)

type Int32Slice []int32

func (arr Int32Slice) Len() int           { return len(arr) }
func (arr Int32Slice) Less(i, j int) bool { return arr[i] < arr[j] }
func (arr Int32Slice) Swap(i, j int)      { arr[i], arr[j] = arr[j], arr[i] }

func main() {

	var test_count int
	fmt.Scanln(&test_count)
	for test_iter := 0; test_iter < test_count; test_iter++ {

		var N, C int
		fmt.Scanln(&N, &C)

		Xs := make([]int32, 0, N)
		for i := 0; i < N; i++ {

			var xi int32
			fmt.Scanln(&xi)
			Xs = append(Xs, xi)
		}

		sort.Sort(Int32Slice(Xs))

		maxdist := (Xs[N-1]-Xs[0])/int32(C-1) + 1
		mindist := int32(0)

		for mindist+1 < maxdist {
			middist := (mindist + maxdist) / 2
			if match(middist, Xs, C-1) {
				mindist = middist
			} else {
				maxdist = middist
			}
		}

		fmt.Println(mindist)
	}
}

func match(dist int32, stalls []int32, cows int) bool {

	nextpos := stalls[0] + dist
	for _, pos := range stalls {
		if pos >= nextpos {
			nextpos = pos + dist
			cows--
			if cows <= 0 {
				return true
			}
		}
	}
	return false
}
