# Lab 1: Imperative Programming in Go

> If you're stuck look at examples on [Go by Example](https://gobyexample.com/)

## Question 1 - Hello World

Below is a complete 'Hello World' program written in Go:

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello World")
}
```

### Question 1a

Type the above program into a new file `hello.go` (don't just copy and paste). To run your code you can either use `go run hello.go` or `go build hello.go` followed by `./hello`. Verify that Hello World is printed in both cases.

### Question 1b

Modify `hello.go` so that it uses a for loop to print Hello World 20 times.

## Question 2 - Fizz Buzz

Write a program that prints the numbers from 1 to 100. But for multiples of 3 print `Fizz` instead of the number and for the multiples of 5 print `Buzz`. For numbers which are multiples of both 3 and 5 print `FizzBuzz`.

## Question 3 - Quiz

Open `quiz.go`. It's a skeleton for a quiz program. Write a `main()` using the provided helper functions so that your program asks the 6 questions from `quiz-questions.csv` and prints out the final score at the end.

<details>
    <summary>Hint 1</summary>

Use `s := score(0)` to initialise your score variable.

</details>

<details>
    <summary>Hint 2</summary>

Use a [for-range loop](https://gobyexample.com/range) to ask all the questions.

</details>

## Question 4 - Arrays vs Slices

Open `sequences.go`.

### Question 4a

Implement `mapSlice` and `mapArray` using [for-range loops](https://gobyexample.com/range).

They are the same as Haskell's map. For example mapping `addOne` onto `[5, 10, 15]` should return `[6, 11, 16]`.

### Question 4b

In `main()`:

- Create a slice `intsSlice` with values `[1, 2, 3]`.
- Map `addOne` onto this slice.
- Print `intsSlice`.

- Create an array `intsArray` of length 3 with values `[1, 2, 3]`.
- Map `addOne` onto this array.
- Print `intsArray`.

Explain the result. Modify the skeleton to solve any issues that you may have observed.

<details>
    <summary>Hint</summary>

Are you passing arguments by value, pointer or reference?
What are the things that

</details>

### Question 4c

Change the definitions of `intsArray` and `intsSlice` so that they both contain values `[1, 2, 3, 4, 5]`. Without modifying `mapSlice` or `mapArray` run your new program. Explain the result.

### Question 4d

Slices support a “slice” operator with the syntax `slice[lowIndex:highIndex]`. It allows you to cut out a portion of your slice. For example:

```go
// Given: intsSlice  = [2, 3, 4, 5, 6]
newSlice := intsSlice[1:3]  
// newSlice = [3, 4]
```

Define `newSlice` as shown above, map `square` onto `newSlice` and print both `newSlice` and the original `intSlice`. Explain the result.

### Question 4e

The function `double` should append a slice onto itself. For example, given `[5, 6, 7]` it should return `[5, 6, 7, 5, 6, 7]`. In `main`, try doubling and then printing your `intsSlice`. Explain the result. Modify the skeleton to solve any issues that you may have observed.

### Question 4f

Given the differences that you found between arrays and slices:

- Explain how arrays and slices are passed to functions.
- Explain how `append()` works.
- Give use cases for arrays and slices.

## Question 5 - Concurrent Hello World

A goroutine is a lightweight thread of execution. Modify your `hello.go` so that it uses a for loop to start 5 goroutines and print `Hello from goroutine i` where `i` is the number of the goroutine.

Example output:

```bash
$ go run hello.go

Hello from goroutine 2
Hello from goroutine 3
Hello from goroutine 4
Hello from goroutine 0
Hello from goroutine 1
```

<details>
    <summary>Hint 1 - How do I start a new goroutine?</summary>

Starting a goroutine is easy, just say:

```go
go someFunc()
```

</details>

<details>
    <summary>Hint 2 - Why does my program exit without printing anything?</summary>

You may notice that your program exits without printing anything. For now you can fix this by placing this after your for loop:

```go
time.Sleep(1 * time.Second)
```

Soon you'll see how to fix this problem with channels.

</details>
