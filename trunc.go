package main

import (
	"fmt"
	"math"
)

func main() {
	var num float64
	fmt.Println("Type a float")
	fmt.Scan(&num)
	num2 := int(math.Trunc(num))
	fmt.Println("Your truncated number is", num2)

}