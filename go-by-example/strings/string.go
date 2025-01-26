package main

import "fmt"

type status int

const (
	Yes status = iota
	No
)

var mapping = map[status]string{
	Yes: "✅",
	No:  "❌",
}

type currentState struct {
	state status
}

func (c *currentState) changeStatus(new status) {
	c.state = new
}

func main() {
	p := currentState{state: No}
	p.changeStatus(No)

	fmt.Println(mapping[p.state])
}
