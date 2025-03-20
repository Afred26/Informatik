package dhbw

import "fmt"

func ExampleFooS() {
	t1 := "abc"
	t2 := "a"
	t3 := "NNM"
	fmt.Println(FooS(t1))
	fmt.Println(FooS(t2))
	fmt.Println(FooS(t3))
	fmt.Println(FooR(t1))
	fmt.Println(FooR(t2))
	fmt.Println(FooR(t3))

	// Output:
	// abcabcabc
	// a
	// NNMNNMNNM
	// abcabcabc
	// a
	// NNMNNMNNM
}
