package main

import (
	"fmt"
)

type Email[T any] struct {
	sender   T
	reciever T
	message  T
}

func emailSender[T any](email <-chan Email[T], done chan bool) {
	for emails := range email {
		fmt.Println("email sent from 📧", emails.sender, " to ", emails.reciever)
	}
	done <- true
}

func main() {
	email := make(chan Email[int], 100)
	done := make(chan bool)

	fmt.Println("Emails started to process")
	go emailSender(email, done)

	for i := 0; i < 10; i++ {
		email <- Email[int]{i, i + 1, i + 4}
	}
	close(email)
	fmt.Println("Done sending ✅")
	<-done
	fmt.Println("Successfull run ✅")

}
