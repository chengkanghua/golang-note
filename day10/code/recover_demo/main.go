package main

import (
	"fmt"
	"sync"
)

func f1() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%v\n", r)
		}
	}()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("recover in goroutine:%v\n", r)
			}
		}()
		defer wg.Done()

		s1 := []int{1, 2}
		fmt.Println(s1[3]) // panic
		// panic("xx")
	}()

	go f2()

	wg.Wait()
	fmt.Println("done")
}

func f2() {
	// ...
}

func main() {
	f1()
}
