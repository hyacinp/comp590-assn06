package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	res   int = 1
	mutex sync.Mutex
)

func worker(n int) {
	mutex.Lock()
	res *= n
	mutex.Unlock()
}

func power(n, p int) {
	for i := 0; i < p; i++ {
		go worker(n)
	}
}

func main() {
	n := 2
	p := 8

	power(n, p)

	time.Sleep(3 * time.Second)

	fmt.Println("Result:", res)
}
