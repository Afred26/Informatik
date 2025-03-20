package aufgabe3

import "fmt"

func ExamplePrime() {
	fmt.Println(Prime(1))
	fmt.Println(Prime(5))
	fmt.Println(Prime(120))
	fmt.Println(Prime(100))
	fmt.Println(Prime(9999))

	//Output:
	//2
	//11
	//659
	//541
	//104723

}
