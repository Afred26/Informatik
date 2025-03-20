package dhbwl

import "fmt"

func ExampleFoo() {
	t1 := "abc"
	t2 := "a"
	t3 := "NNM"
	fmt.Println(Foo(t1, len(t1)))
	fmt.Println(Foo(t2, len(t2)))
	fmt.Println(Foo(t3, len(t3)))
	// Output:
	// abcabcabc
	// a
	// NNMNNMNNM
}
