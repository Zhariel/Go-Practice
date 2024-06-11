package main

import (
	"fmt"
	"sync"
)

func read(x *int, wg *sync.WaitGroup){
	fmt.Println(*x)
	wg.Done()
}

func add(x *int, wg *sync.WaitGroup){
	*x += 1
	fmt.Println(*x)
	wg.Done()
}

func main() {
	var wg sync.WaitGroup
	x := 1
	wg.Add(2)
	go read(&x, &wg)
	go add(&x, &wg)
	wg.Wait()
}

//Read and Add both access the x variable.
//Add increments and prints so it always prints 2.
//Read just prints the variable so it should print 1 but in practice can print both 1 and 2.
//This produces non deterministic results, and therefore is a very bad oopsie and a big no no.