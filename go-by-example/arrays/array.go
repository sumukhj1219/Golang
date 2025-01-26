package main

import "fmt"

func example(nums ...int) {
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Sum is = ", total)
}

func main() {
	nums := []int{1, 3, 42, 1}
	example(nums...)
}
