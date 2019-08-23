package main

import "fmt"

func say(something string, done chan<- bool) {
	fmt.Println(something)
	done <- true
}

func main() {
	done := make(chan bool)
	go say("Hello " + "World", done)
	<-done
}