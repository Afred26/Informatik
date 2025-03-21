package Aufgabe1

import (
	"fmt"
)

func ExampleMult() {
	v1 := []float64{23.0, 1.5, 12.0, 31.0}
	v2 := []float64{12.42, 11.01, 10.12, 9.24, 8.34}
	v3 := []int{3, 29, 4, 0, 42}
	v4 := []int{}

	fmt.Println(Mult(v1, v3))
	fmt.Println(Mult(v2, v4))
	fmt.Println(Mult(v2, v3))

	// Output:
	// [69 43.5 48 0]
	// []
	// [37.26 319.29 40.48 0 350.28]
}
