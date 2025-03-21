package Aufgabe3

import "fmt"

func ExampleIsPrim() {
	v := 15
	w := 7
	x := 43
	y := 1
	z := -23
	fmt.Println(IsPrim(v))
	fmt.Println(IsPrim(w))
	fmt.Println(IsPrim(x))
	fmt.Println(IsPrim(y))
	fmt.Println(IsPrim(z))

	// Output:
	//
	//false
	//true
	//true
	//false
	//false
}
