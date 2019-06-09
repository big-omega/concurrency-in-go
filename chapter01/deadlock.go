package main

import (
	"fmt"
	"sync"
	"time"
)

type Value struct {
	m   sync.Mutex
	val int
}

func main() {
	var wg sync.WaitGroup

	printSum := func(v1, v2 *Value) {
		defer wg.Done()

		v1.m.Lock()
		defer v1.m.Unlock()

		time.Sleep(2 * time.Second)

		v2.m.Lock()
		defer v2.m.Unlock()

		fmt.Printf("sum = %d\n", v1.val+v2.val)
	}

	v1 := &Value{sync.Mutex{}, 1}
	v2 := &Value{sync.Mutex{}, 2}
	wg.Add(2)
	go printSum(v1, v2)
	go printSum(v2, v1)
	wg.Wait()
}
