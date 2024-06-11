package main

import (
	"fmt"
	"strings"
)

func check(s string) string {
	if strings.HasPrefix(s, "i") && strings.Contains(s, "a") && strings.HasSuffix(s, "n") {
		return "Found!"
	}
	return "Not Found!"
}

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(check(strings.ToLower(s)))
}
