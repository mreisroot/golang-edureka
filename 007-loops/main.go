package main

import "fmt"

func main() {
	// For loop
	fmt.Printf("\nNormal Go for loop\n")
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}

	// Implementation of a while loop
	fmt.Printf("\nImplementation of a while loop\n")
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// Loop inside loop
	fmt.Printf("\nLoop inside loop")
	for i := 1; i < 5; i++ {
		for j := 1; j < i; j++ {
			fmt.Println("*")
		}
		fmt.Println()
	}

}
