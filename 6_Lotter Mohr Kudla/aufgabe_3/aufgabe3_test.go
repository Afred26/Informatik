package aufgabe3

import "fmt"

func ExampleAddEachElement() {

	l1 := []int{1, 2, 3, 4, 5}
	l2 := []int{8, 9, 7, 6, 5}
	l3 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	l4 := []int{}

	fmt.Println(AddEachElement(l1, l2))
	fmt.Println(AddEachElement(l1, l4))
	fmt.Println(AddEachElement(l1, l3))
	fmt.Println(AddEachElement(l2, l4))
	fmt.Println(AddEachElement(l4, l4))

	// Output:
	// [9 11 10 10 10]
	// [1 2 3 4 5]
	// [2 4 6 8 10 11 12 13 14 15]
	// [8 9 7 6 5]
	// []
}
