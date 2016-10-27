package main

import (
	"fmt"
)

func main() {

	var N int
	fmt.Scanln(&N)

	stds := make([][]bool, 0, N)

	for i := 0; i < N; i++ {

		std := make([]bool, 0, N)
		for j := 0; j < N; j++ {

			var stat string
			fmt.Scan(&stat)
			std = append(std, stat == "Y")
		}
		stds = append(stds, std)
	}


}
