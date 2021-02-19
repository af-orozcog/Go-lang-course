package main

import "fmt"

func GenDisplaceFn(a, s_0, v_0 float64) func(t float64) float64 {
	re := func(t float64) float64 {
		return (0.5 * a * t * t) + (v_0 * t) + s_0
	}
	return re
}

func main() {
	fmt.Println("Please enter the acceleration")
	var a, v, s, t float64
	fmt.Scan(&a)
	fmt.Println("Please enter the initial velocity")
	fmt.Scan(&v)
	fmt.Println("Please enter the intial displacement")
	fmt.Scan(&s)
	fn := GenDisplaceFn(a, s, v)
	fmt.Println("Please enter the time for which you want to know the total displacement")
	fmt.Scan(&t)
	fmt.Printf("the total displacement is ")
	fmt.Println(fn(t))
}
