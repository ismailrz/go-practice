package main

import (
	"fmt"
	"sync"
)

func task(wg *sync.WaitGroup, name string) {

	fmt.Printf("Task %v\n", name)
	wg.Done()

}

func main() {

	var wg sync.WaitGroup

	count := 0

	for i := 0; i < 100; i++ {

		wg.Add(1)
		go func() {
			count++
			fmt.Println(count)
		}()
	}

	wg.Add(1)
	go task(&wg, "Login Page")

	go task(&wg, "Signup Page")

	wg.Wait()

	// time.Sleep(time.Second * 6)
}
