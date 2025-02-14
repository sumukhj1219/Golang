package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println("Error in opening env file")
	}
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error in reading file", err)
	}
}
