package main

import (
	"fmt"
	"time"
)

type Task[T int] struct {
	num T
}

type Result[T int] struct {
	num T
}

func worker[T int](taskChan <-chan Task[T], resultChan chan<- Result[T], done chan<- bool) {
	fmt.Println("Scheduling the tasks to results ⏲️")
	for tasks := range taskChan {
		resultChan <- Result[T]{tasks.num * tasks.num}
		fmt.Println("Scheduled task ", tasks.num, "at time :", time.TimeOnly)
	}
	done <- true
}

func main() {
	taskChan := make(chan Task[int], 50)
	resultChan := make(chan Result[int], 50)
	done := make(chan bool)

	go worker(taskChan, resultChan, done)

	for i := 0; i < 10; i++ {
		taskChan <- Task[int]{num: i}
	}
	close(taskChan)
	fmt.Println("Loaded all tasks in tasks scheduler ✅")
	<-done
	fmt.Println("Success ✅")

}
