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
	messages := make(chan string)

	// Start a new goroutine that will send some messages.
	go sendMessages(messages)

	// Receive the 3 messages sent by the goroutine.
	for i := 0; i < 3; i++ {
		// Wait 1s between each receive.
		time.Sleep(1 * time.Second)
		receivedMessage := <-messages
		fmt.Println("Main has received:", receivedMessage)
	}
}
