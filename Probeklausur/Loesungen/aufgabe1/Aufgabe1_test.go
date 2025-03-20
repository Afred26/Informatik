package aufgabe1

import "fmt"

func ExampleFoo() {

	fmt.Println(Foo(1, 1, 1))
	fmt.Println(Foo(1, 1, 5))
	fmt.Println(Foo(1, 1, 30))
	fmt.Println(Foo(2, 1, 6))

	//Output:
	// 1
	// 5
	// 832040
	// 11
}
