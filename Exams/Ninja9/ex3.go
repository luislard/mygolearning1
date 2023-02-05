package main

import (
	"fmt"
	"runtime"
	"sync"
)

/**
Create something with race condition
*/
func main() {

	counter := 0

	const gs = 100
	var wg sync.WaitGroup
	wg.Add(gs)

	for i := 0; i < gs; i++ {
		go func() {
			v := counter
			runtime.Gosched()
			v++
			counter = v
			wg.Done()
		}()
		fmt.Println("go routines", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("count:", counter)

}
