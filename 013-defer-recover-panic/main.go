package main

import "fmt"

func main() {
	// The functions will be executed in order
	FirstRun()
	SecondRun()

	fmt.Println()

	// SecondRun will be executed first
	defer FirstRun()
	SecondRun()

	fmt.Println()

	// Error handling with recover and panic
	fmt.Println(div(3, 0))
	fmt.Println(div(3, 5))
	demPanic()
}

func FirstRun() {
	fmt.Println("I executed First")
}

func SecondRun() {
	fmt.Println("I executed Second")
}

func div(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()

	solution := num1 / num2
	return solution
}

func demPanic() {
	defer func() {
		fmt.Println(recover())
	}()

	panic("Panic")
}
