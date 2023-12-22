package main

import "fmt"

func main() {
	/* Declaration of a map in which
	   the key is of type string
	   and the value is of type int */
	StudentAge := make(map[string]int)

	/* The first example is that
	   Aryya is the key and
	   23 is the value */
	StudentAge["Aryya"] = 23
	StudentAge["Saurabh"] = 27
	StudentAge["Prerna"] = 21
	StudentAge["Akrati"] = 19
	StudentAge["Sahit"] = 42
	StudentAge["Kirti"] = 22

	// Print the entire map
	fmt.Println(StudentAge)

	fmt.Println()

	// Print the value of a specific key in the map
	fmt.Println(StudentAge["Kirti"])

	fmt.Println()

	// Print the length of the map
	fmt.Println(len(StudentAge))

	fmt.Println()

	// Maps inside of maps
	superhero := map[string]map[string]string{

		"Superman": map[string]string{
			"RealName": "Clark Kent",
			"City":     "Metropolis",
		},

		"Batman": map[string]string{
			"RealName": "Bruce Wayne",
			"City":     "Gotham",
		},
	}

	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["RealName"], temp["City"])
	}

}
