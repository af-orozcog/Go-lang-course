package main

import "fmt"

func main() {
	var value float32

	fmt.Scan(&value)

	var valueInt int = int(value)

	fmt.Printf("%d\n", valueInt)
}
