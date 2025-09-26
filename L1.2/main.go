package main

import (
	"fmt"
	"sync"
)

func main() {
	arr := [5]int{2, 4, 6, 8, 10}
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, num := range arr {
		wg.Add(1)
		go func(x int) {
			defer wg.Done()
			square := x * x

			mu.Lock()
			fmt.Println(square)
			mu.Unlock()
		}(num)
	}

	wg.Wait()
}
