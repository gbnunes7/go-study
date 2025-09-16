// package main

// import (
// 	"fmt"
// 	"time"
// )

// func showMessage() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Hello from goroutine2:", i)
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func sayHello() {
// 	for i := 0; i < 5; i++ {
// 		fmt.Println("Hello from goroutine:", i)
// 		time.Sleep(100 * time.Millisecond)
// 	}
// }

// func main() {
// 	go sayHello()
// 	go showMessage()
// 	time.Sleep(1 * time.Second)
// }