package dhbwl

import "fmt"

func ExampleFoo() {
	t1 := "abc"
	t2 := "a"
	t3 := "NNM"
	fmt.Println(Foo(t1, t1))
	fmt.Println(Foo(t2, t2))
	fmt.Println(Foo(t3, t3))
	// Output:
	// abcabcabc
	// a
	// NNMNNMNNM
}
