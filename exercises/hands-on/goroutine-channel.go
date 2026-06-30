package main

import "fmt"

func worker(ch chan string, name string) {

	ch <- name
}

func main() {

	ch := make(chan string, 3)

	go worker(ch, "Hasan Sheikh")

	message := <-ch

	fmt.Println(message)

}
