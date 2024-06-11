package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

func main() {
	fmt.Println("Enter your ints separated by spaces.")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	sli := strings.Split(strings.TrimSpace(line), " ")
	var nums []int
	for _, s := range sli {
		num, _ := strconv.Atoi(s)
		nums = append(nums, num)
	}

	var wg sync.WaitGroup
	var mu sync.Mutex
	length := len(nums)
	if length < 4 {
		fmt.Println("Array needs to contain at least 4 values.")
		os.Exit(1)
	}

	l := length / 4
	var concat_slice []int
	for n := 0; n < 4; n++ {
		start := n * l
		end := start + l
		if n == 3 {
			end = length
		}

		wg.Add(1)
		go func(sli []int, concat *[]int) {
			defer wg.Done()
			sort.Ints(sli)
			fmt.Println(sli)
			mu.Lock()
			defer mu.Unlock()
			*concat = append(*concat, sli...)
		}(nums[start:end], &concat_slice)
	}

	wg.Wait()
	fmt.Println("Merged final: ", concat_slice)
	sort.Ints(concat_slice)
	fmt.Println("Sorted final: ", concat_slice)
}
