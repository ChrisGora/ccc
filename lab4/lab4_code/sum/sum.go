package main

import (
	"fmt"
	"sync"
)

func main() {
	sum := 0

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			sum = sum + 1
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println(sum)
}
