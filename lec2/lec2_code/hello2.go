package main

import (
	"fmt"
	"time"
)

func say(something string) {
	fmt.Println(something)
}

func main() {
	go say("Hello " + "World")
	time.Sleep(1 * time.Second)
}