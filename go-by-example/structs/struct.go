package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := new(Person)
	person.Name = "sumukh"
	person.Age = 21
	fmt.Println(*person)
}
