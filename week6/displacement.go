package main

import (
	"fmt"
	"math"
)

func GetDisplaceFn(a, v0, s0 float64) func(float64) float64 {
	return func(t float64) float64 {
		return 0.5*a*math.Pow(t, 2) + v0*t + s0
	}
}

func main() {
	var a, v0, s0 float64
	fmt.Println("enter acceleration")
	fmt.Scan(&a)

	fmt.Println("enter initial speed")
	fmt.Scan(&v0)

	fmt.Println("enter initial displacement")
	fmt.Scan(&s0)

	f := GetDisplaceFn(a, v0, s0)

	var t float64
	fmt.Println("enter time")
	fmt.Scan(&t)

	fmt.Println(f(t))
}
