package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {
	s := make([]int, 0)
	var x string

	fmt.Println("Input your numbers.")
	for {
		fmt.Scan(&x)
		if strings.ToLower(x) == "x" {
			break
		}

		if elt, err := strconv.Atoi(x); err == nil {
			s = append(s, elt)
		}
	}
	sort.Ints(s)
	fmt.Println(s)
}
