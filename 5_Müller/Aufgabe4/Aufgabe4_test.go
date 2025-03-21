package Aufgabe4

import "fmt"

func ExampleFilterNumbers() {
	v := "123"
	x := "Hallo0W1eG3ht2?"
	y := "1A234567B890C"
	z := "ABC"
	fmt.Println(FilterNumbers(v))
	fmt.Println(FilterNumbers(x))
	fmt.Println(FilterNumbers(y))
	fmt.Println(FilterNumbers(z))

	// Output:
	//
	//HalloWeGht?
	//ABC
	//ABC
}
