package main

import "fmt"

func main() {
	var name string = "Aryya Paul"
	fmt.Println(len(name))
	fmt.Println(name + " is a chill dude")

	const pi float64 = 3.1412435
	x := 5
	isbool := true
	win := true

	fmt.Printf("\n%.3f \n", pi)
	fmt.Printf("%T \n", name)
	fmt.Printf("%T \n", isbool)
	fmt.Printf("%t \n", pi)
	fmt.Printf("%t \n", win)
	fmt.Printf("%d \n", x)
	fmt.Printf("%b \n", 25)
	fmt.Printf("%c \n", 34)
	fmt.Printf("%x \n", 15)
	fmt.Printf("%e \n", pi)
}
