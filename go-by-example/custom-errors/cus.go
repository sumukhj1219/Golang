package main

import "fmt"

type customErrors[T string] struct {
	arg     T
	message T
}

func (c *customErrors[T]) Error() T {
	return T(fmt.Sprintf("%s - %s", c.arg, c.message))
}

func f(arg int) (int, error) {
	if arg == 42 {
		return -1, &customErrors[string]{arg: "caught", message: "42 cannot be sent"}
	}
	return 1, nil
}

func main() {
	result, err := f(42)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
