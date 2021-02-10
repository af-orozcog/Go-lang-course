package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	var vec []int = make([]int, 0, 3)

	for {
		var temp string

		fmt.Scan(&temp)

		if temp == "X" {
			break
		}

		number, err := strconv.Atoi(temp)

		_ = err

		vec = append(vec, number)

		sort.Ints(vec)

		for i, v := range vec {
			fmt.Printf("%d ", v)
			i = i + 1
		}
		fmt.Print("\n")
	}
}
