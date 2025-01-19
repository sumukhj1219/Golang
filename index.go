package main

import "fmt"

func utility(name string) string {
	var message string = fmt.Sprintf("Hi I am %v", name)
	return message
}

func main() {
	// fmt.Println(utility("Sumukh"))

	// var o [5]int = [5]int{1, 2, 3, 4, 5}

	p := []int{1, 2, 6, 7, 0}
	p = append(p, 2)
	fmt.Println("p:=", p)

}
