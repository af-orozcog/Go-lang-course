package main

import "fmt"

func swap(arr []int, ind1 int, ind2 int) {
	temp := arr[ind1]
	arr[ind1] = ind2
	arr[ind2] = temp
}

func BubbleSort(arr []int) {
	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			if arr[j] > arr[j+1] {
				swap(arr, j, j+1)
			}
		}
	}
}

func main() {
	for {
		var n int

		fmt.Println("Please type the size of the array: ")
		fmt.Scan(&n)

		if n <= 0 {
			fmt.Println("Good bye!")
			break
		}

		arr := make([]int, n)

		for i := 0; i < n; i++ {
			fmt.Scan(&arr[i])
		}

		BubbleSort(arr)

		for _, v := range arr {
			fmt.Printf("%d ", v)
		}
		fmt.Println()
	}
}
