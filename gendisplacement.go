package main

import (
	"fmt"
)

func GenDisplaceFn(a, vo, so float64) func(float64) float64 {
	return func(t float64) float64{
		return 0.5 * a * t * t + vo * t + so
	}
}

func main() {
	var a, vo, so, t float64

	fmt.Print("Input acceleration: ")
	fmt.Scan(&a)
	
	fmt.Print("Input initial velocity: ")
	fmt.Scan(&vo)
	
	fmt.Print("Input initial displacement: ")
	fmt.Scan(&so)
	
	fmt.Print("Input time: ")
	fmt.Scan(&t)

	fn := GenDisplaceFn(a, vo, so)
	fmt.Println(fn(t))

}
