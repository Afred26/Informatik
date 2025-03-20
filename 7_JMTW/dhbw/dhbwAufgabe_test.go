package dhbw

import "fmt"

func ExampleFoo() {
	t1 := "abc"
	t2 := "a"
	t3 := "NNMNNMNNM"
	fmt.Println(Foo(t1))
	fmt.Println(Foo(t2))
	fmt.Println(Foo(t3))

	// Output:
	// abcabcabc
	// a
	// NNMNNMNNM
}
