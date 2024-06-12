package main

import (
	"fmt"
	"sync"
)

type Chop struct {
	sync.Mutex
}

type Philo struct {
	lchop, rchop *Chop
	id           int
}

func (p Philo) eat(perms chan []int, eatCount []int, wg2 *sync.WaitGroup) {
	defer wg2.Done()
	perm := <-perms
	if (p.id != perm[0] && p.id != perm[1]) || eatCount[p.id] <= 0 {
		return
	}
	p.lchop.Lock()
	p.rchop.Lock()
	fmt.Println("Eating", p.id)
	eatCount[p.id]--
	p.lchop.Unlock()
	p.rchop.Unlock()
	fmt.Println("Finished eating", p.id)
}

func sum(sli []int) int {
	n := 0
	for _, v := range sli {
		n += v
	}
	return n
}

func main() {
	var wg sync.WaitGroup
	var wg2 sync.WaitGroup

	sticks := make([]Chop, 5)
	philos := make([]Philo, 5)
	for i := range philos {
		philos[i] = Philo{&sticks[i], &sticks[(i+1)%5], i}
	}
	eatCount := []int{3, 3, 3, 3, 3}
	perms := make(chan []int, 5)

	wg.Add(1)
	go func() { //host
		defer wg.Done()
		x := 0
		for sum(eatCount) > 0 {
			selected := []int{x, (x + 2) % 5}
			for i := 0; i < 5; i++ {
				perms <- selected
				wg2.Add(1)
				go philos[i].eat(perms, eatCount, &wg2)
			}
			wg2.Wait()
			x = (x + 1) % 5
		}
	}()

	wg.Wait()
}
