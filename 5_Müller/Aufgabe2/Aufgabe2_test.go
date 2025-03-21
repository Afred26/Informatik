package Aufgabe2

import "fmt"

func ExampleCountVocals() {
	s1 := "Hallo, es ist ein schöner Tag."
	s2 := "Wir haben Donnerstag und machen Informatik."
	s3 := ""

	fmt.Println(CountVocals(s1))
	fmt.Println(CountVocals(s2))
	fmt.Println(CountVocals(s3))

	// Output:
	// [2 3 2 1 0]
	// [4 3 3 2 1]
	// [0 0 0 0 0]
}
