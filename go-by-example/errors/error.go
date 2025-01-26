package main

import (
	"errors"
	"fmt"
)

func f[T int](arg T) (T, error) {
	if arg == 42 {
		return -1, errors.New("Cant cook right now as 42")
	}
	return arg + 3, nil
}

var ErrOutOfSupplies error = fmt.Errorf("Out of supplies")
var ErrNoPower error = fmt.Errorf("Power went off it's quite dark")

func makeTea[T int](arg T) (T, error) {
	if arg == 2 {
		return -1, ErrOutOfSupplies
	}
	if arg == 3 {
		return -1, ErrNoPower
	} else {
		return arg, nil
	}
}

func main() {
	for _, val := range []int{7, 42} {
		if res, err := f(val); err != nil {
			fmt.Println("failed", err)
		} else {
			fmt.Println("success", res)
		}
	}
}
