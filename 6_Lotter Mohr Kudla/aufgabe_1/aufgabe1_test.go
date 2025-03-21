package aufgabe1

import "fmt"

func ExampleStringsWithoutDigit() {
	l1 := []string{""}
	l2 := []string{"abc", "123", "a1b2c3", "hello"}
	l3 := []string{"123", "456", "789"}
	l4 := []string{"noDigitsHere", "anotherOne", "123abc", "pureText"}

	fmt.Println(StringsWithoutDigit(l1))
	fmt.Println(StringsWithoutDigit(l2))
	fmt.Println(StringsWithoutDigit(l3))
	fmt.Println(StringsWithoutDigit(l4))

	// Output:
	// []
	// [abc hello]
	// []
	// [noDigitsHere anotherOne pureText]
}
