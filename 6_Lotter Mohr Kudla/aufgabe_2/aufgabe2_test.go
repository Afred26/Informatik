package aufgabe2

import (
	"fmt"
)

func ExampleEverythingBevorSubstring() {
	fmt.Println(EverythingBevorSubstring([]string{"abc123", "def456", "ghi789"}, "123"))
	fmt.Println(EverythingBevorSubstring([]string{"prefix_test", "another_prefix", "noaffix"}, "prefix"))
	fmt.Println(EverythingBevorSubstring([]string{"hello_world", "world_hello", "test"}, "world"))
	fmt.Println(EverythingBevorSubstring([]string{}, "empty"))
	fmt.Println(EverythingBevorSubstring([]string{"noaffixhere", "stillnoaffix"}, "affix"))

	// Output:
	// [abc def456 ghi789]
	// [another_ noaffix]
	// [hello_ test]
	// [empty]
	// [no stillno]
}
