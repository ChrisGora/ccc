package main

import "fmt"

func goodbye(array []string) {
	array[0] = "goodbye"
}

func main() {
	//s := make([]string, 2)
	//s[0] = "Hello"
	//s[1] = "World"
	//fmt.Println(s)

	s := []string{"Hello", "World"}
	fmt.Println(s)

	goodbye(s)

	fmt.Println(s)
}
