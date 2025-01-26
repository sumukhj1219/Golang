package main

import "fmt"

func main() {
	maps := make(map[string]int)
	maps["apple"] = 1

	nestedMaps := make(map[int]map[string]int)
	nestedMaps[2]["apple"] = 1

	fmt.Println(nestedMaps)
	fmt.Println(maps)
}
