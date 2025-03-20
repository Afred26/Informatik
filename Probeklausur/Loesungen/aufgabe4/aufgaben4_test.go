package aufgabe4

import "fmt"

func Example() {
	x1 := Pythagoreantriplet{a: 3, b: 4, c: 5}
	x2 := Pythagoreantriplet{a: 3, b: 4, c: 9}
	x3 := Pythagoreantriplet{a: 4, b: 3, c: 5}
	x4 := Pythagoreantriplet{a: 0, b: 500, c: 500}
	fmt.Println(x1.IsPythagorean())
	fmt.Println(x2.IsPythagorean())
	fmt.Println(x3.IsPythagorean())
	fmt.Println(x4.IsPythagorean())
	fmt.Println(FindTriplets(1000))

	// Output:
	// true
	// false
	// false
	// false
	// {200 375 425}
}
