package main

import "fmt"

type Options int

const (
	Yes Options = iota
	No
)

type Items struct {
	Id        int
	Name      string
	CreatedOn string
	Status    Options
}

type Todos struct {
	items    []Items
	globalId int
}

type TodoInterface interface {
	addTodo(name string, createdOn string, status int)
	deleteTodo(id int)
	getTodos() []Items
}

func (t *Todos) addTodo(name string, createdOn string, status int) {
	t.globalId = t.globalId + 1

	newItems := Items{
		Id:        t.globalId,
		Name:      name,
		CreatedOn: createdOn,
		Status:    Options(status),
	}
	t.items = append(t.items, newItems)
	fmt.Println("New todos added successfully ✅", newItems)
}

func (t *Todos) deleteTodo(id int) {
	var filteredTodos []Items
	for _, value := range t.items {
		if value.Id != id {
			filteredTodos = append(filteredTodos, value)
		}
	}
	t.items = filteredTodos
	fmt.Println("Todos after deleting :", t.items)
}

func (t *Todos) getTodos() {
	for _, value := range t.items {
		fmt.Println("\n", value)
	}
}

func main() {

	todoList := &Todos{}
	todoList.addTodo("milk", "21/01/2025", 1)
	todoList.addTodo("milk", "21/01/2025", 2)
	todoList.addTodo("milk", "21/01/2025", 4)
	todoList.addTodo("milk", "21/01/2025", 1)

	todoList.deleteTodo(1)
	todoList.getTodos()
}
