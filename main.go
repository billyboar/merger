package main

import (
	"fmt"
	"math/rand"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4, 6, 8)
	c := merge(a, b)
	wg.Add(2)
	for v := range c {
		fmt.Println(v)
	}
	/*
	   this should print something like this:
	       1
	       2
	       3
	       4
	       5
	       6
	       7
	       8
	   make sure that you close both channels and program should exit successfully at the end.
	*/
}

func merge(a, b <-chan int) <-chan int {
	// this should take a and b and return a new channel which will send all values from both a and b
	ch := make(chan int)
	go func() {
		for num := range a {
			ch <- num
		}
		wg.Done()
	}()

	go func() {
		for num := range b {
			ch <- num
		}
		wg.Done()
	}()

	go func() {
		wg.Wait()
		close(ch)
	}()
	return ch
}

func asChan(vs ...int) <-chan int {
	// this should reutrn a channel and send `vs` values randomly to it.
	for i := range vs {
		l := rand.Intn(i + 1)
		vs[i], vs[l] = vs[l], vs[i]
	}
	ch := make(chan int)
	go func() {
		for _, v := range vs {
			ch <- v
		}
		close(ch)
	}()
	return ch
}
