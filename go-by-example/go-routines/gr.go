package main

import (
	"fmt"
)

func processNum(num chan int, a, b int) {
	sum := a + b
	num <- sum // Send the sum to the channel
}

func main() {
	num := make(chan int) // Create a channel for integers

	go processNum(num, 4, 5) // Call processNum in a goroutine

	res := <-num // Receive the result from the channel

	fmt.Println("Sum =", res) // Print the result
}
