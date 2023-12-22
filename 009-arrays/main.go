package main

import "fmt"

func main() {
	// First array
	var StudentsCount [10]int

	for i := 0; i < 10; i++ {
		StudentsCount[i] = i + 1
		fmt.Println(StudentsCount[i])
	}

	fmt.Println()

	// Second array
	/* var EvenNum [5]int

	EvenNum[0] = 0
	EvenNum[1] = 2
	EvenNum[2] = 4
	EvenNum[3] = 6
	EvenNum[4] = 8

	fmt.Println(EvenNum[2]) */

	fmt.Println()

	// Another way to declare and initialize the second array
	EvenNum := [5]int{0, 2, 4, 6, 8}
	fmt.Println(EvenNum[2])

	// Second way to loop through arrays
	/* The _ means that we're not actually
	   using the iterator and the value is
	   in the range EvenNum */
	for _, value := range EvenNum {
		fmt.Println(value)
	}

	fmt.Println()

	/* If you'd like to use the iterator
	   and print values along with their
	   indexes, you could do */
	for i, value := range EvenNum {
		fmt.Println(value, i)
	}

	fmt.Println()

	// Slicing arrays
	numSlice := []int{5, 4, 3, 2, 1}

	sliceA := numSlice[3:5]
	fmt.Println(sliceA)

	sliceB := numSlice[:5]
	fmt.Println(sliceB)

	sliceC := numSlice[0:]
	fmt.Println(sliceC)

	sliceD := make([]int, 5, 10)
	copy(sliceD, numSlice)

	fmt.Println()

	fmt.Println(sliceD)
	fmt.Println(numSlice)

	sliceE := append(numSlice, 3, 0, -1)
	fmt.Println(sliceE)
}
