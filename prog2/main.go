package main

import (
	"fmt"
	"sync"
)

var (
	res   int = 1
	mutex sync.Mutex
	wg    sync.WaitGroup
)

func worker(n int) {
	defer wg.Done()
	mutex.Lock()
	res *= n
	mutex.Unlock()
}

func power(n, p int) {
	for i := 0; i < p; i++ {
		wg.Add(1)
		go worker(n)
	}
}

func main() {
	n := 2
	p := 8
	power(n, p)

	wg.Wait()

	fmt.Println("Result:", res)
}
