package main

import (
	"fmt"
	"time"
)

func power_h(n, p, res int) {
	if p == 1 {
		fmt.Println("Result:", res)
		return
	}
	go power_h(n, p-1, res*n)
}

func power(n, p int) {
	go power_h(n, p, n)
}

func main() {
	n := 2
	p := 3

	power(n, p)

	time.Sleep(3 * time.Second)

}
