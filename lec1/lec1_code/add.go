package main

import "fmt"

func add(a int, b int) int {
	sum := a + b
	return sum
}

func main() {
	firstNumber := 10
	secondNumber := 5
	firstNumber = add(firstNumber, secondNumber)
	fmt.Println(firstNumber)
}
