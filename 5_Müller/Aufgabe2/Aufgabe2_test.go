package Aufgabe2

import "fmt"

func ExampleCountUmlauts() {
	s1 := "Hallo, es ist ein schöner Tag."
	s2 := "Wir haben Donnerstag und machen Informatik."
	s3 := ""

	fmt.Println(CountUmlauts(s1))
	fmt.Println(CountUmlauts(s2))
	fmt.Println(CountUmlauts(s3))

	// Output:
	// [2 3 2 1 0]
	// [4 3 3 2 1]
	// [0 0 0 0 0]
}
