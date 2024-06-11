package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type person struct {
	name     string
	lastname string
}

func main() {
	var filename string

	fmt.Println("Enter filename: ")
	fmt.Scan(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error reading csv", err)
		return
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = ' '

	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading records", err)
	}

	people_slice := make([]person, 0)
	for _, p := range records {
		pstruct := person{p[0], p[1]}
		people_slice = append(people_slice, pstruct)
	}

	for p := range people_slice {
		fmt.Println(people_slice[p].name, people_slice[p].lastname)
	}
}
