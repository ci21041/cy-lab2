package main

import (
	"fmt"
	"time"
)

func sendMessages(receiver chan string) {
	// Create a slice of some strings to send.
	messages := []string{
		"ping",
		"pong",
		"pinggg",
	}

	// Send the 3 messages to the receiver
	for _, m := range messages {
		fmt.Println("sendMessages is sending:", m)
		receiver <- m

	}
}

func main() {
	// Create a channel for sending and receiving strings.
	//messages := make(chan string)

	// buffer one will only be executed when it has the size of 3
	newMessages := make(chan string, 3)

	// Start a new goroutine that will send some messages.
	go sendMessages(newMessages)

	// Receive the 3 messages sent by the goroutine.
	for i := 0; i < 3; i++ {
		// Wait 1s between each receive.
		time.Sleep(1 * time.Second)
		receivedMessage := <-newMessages
		fmt.Println("Main has received:", receivedMessage)
	}
}
