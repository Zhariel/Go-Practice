package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func swap(sli []int, i int, j int){
	sli[i], sli[j] = sli[j], sli[i]
}

func bubbleSort(sli []int) {
	for i := 0; i < len(sli)-1; i++ {
		for j := i + 1; j < len(sli); j++ {
			if sli[i] > sli[j] {
				swap(sli, i, j)
			}
		}
	}
}

func main() {
	fmt.Println("Enter your ints separated by spaces.")
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	sli := strings.Split(strings.TrimSpace(line), " ")
	numslice := make([]int, 10)
	for i, s := range sli {
		num, _ := strconv.Atoi(s)
		numslice[i] = num
		if i >= 9 {
			break
		}
	}

	bubbleSort(numslice)

	fmt.Println(numslice)
}
