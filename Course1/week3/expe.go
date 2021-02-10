package main

import "fmt"

func main() {
	x := [...]int{1, 2, 3, 4, 5}

	y := x[0:2]

	z := x[1:4]
	fmt.Println(len(y), cap(y), len(z), cap(z))
}
