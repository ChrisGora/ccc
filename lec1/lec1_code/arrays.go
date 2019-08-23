package main

import "fmt"

func goodbye(array [2]string) {
	array[0] = "goodbye"
}

func main() {
	//var a [2]string
	//a[0] = "Hello"
	//a[1] = "World"
	//fmt.Println(a)

	a := [2]string{"Hello", "World"}
	fmt.Println(a)

	goodbye(a)

	fmt.Println(a)
}
