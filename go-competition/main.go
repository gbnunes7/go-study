package main

import (
	"fmt"
	"time"
)

// func main() {
// 	ch := make(chan int, 3)

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			ch <- i
// 		}
// 		close(ch)
// 		fmt.Printf("Producer done\n")
// 	}()

// 	time.Sleep(1 * time.Second)
// 	for val := range ch {
// 		fmt.Printf("Received: %d\n", val)
// 	}
// }

func producer(c chan<- int) {
	for i := 0; i < 5; i++ {
		c <- i
		fmt.Printf("Produced: %d\n", i)
		time.Sleep(100 * time.Millisecond)
	}
	close(c)
}

func consumer(c <-chan int) {
	for val := range c {
		fmt.Printf("Consumed: %d\n", val)
		time.Sleep(150 * time.Millisecond)
	}
}

func main() {
	ch := make(chan int)
	go producer(ch)
	consumer(ch)

}