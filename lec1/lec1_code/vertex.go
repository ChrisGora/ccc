package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	pointer := &v
	pointer.X = 1e9
	fmt.Println(v) // = {1000000000 2}
}