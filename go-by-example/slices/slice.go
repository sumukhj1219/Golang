package main

import (
	"fmt"
	"sort"
)

func main() {
	slice := []int{1, 3, 4, 4, 8, 9, 7, 6}
	fmt.Println(slice)
	slice = append(slice, 5)
	sort.Ints(slice)

	filteredSlice := []int{}

	for _, v := range slice {
		if v != 5 {
			filteredSlice = append(filteredSlice, v)
		}
	}
	fmt.Println(filteredSlice)
}
