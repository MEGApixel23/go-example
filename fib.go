package main

import "fmt"

func fibonacci(c chan<- int, quit <-chan bool) {
	defer close(c)

	x, y := 0, 1

	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println(`Quit`)
			return
		}
	}

	return
}

func main() {
	c := make(chan int)
	quit := make(chan bool)

	go func(quit chan<- bool) {
		defer close(quit)

		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- true
	}(quit)

	fibonacci(c, quit)
}
