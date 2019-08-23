package main

import (
	"fmt"
	"time"
)

func main() {
	timer := time.After(2 * time.Second)
	fmt.Println("Timer started")
	<-timer
	fmt.Println("Timer expired")
}
