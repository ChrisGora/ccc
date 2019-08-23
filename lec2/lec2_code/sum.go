package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	mid := len(s) / 2

	leftSum := make(chan int)
	rightSum := make(chan int)

	go sum(s[:mid], leftSum)
	go sum(s[mid:], rightSum)

	x := <-leftSum
	y := <-rightSum

	fmt.Println(x, "+", y, "=", x+y)
}
