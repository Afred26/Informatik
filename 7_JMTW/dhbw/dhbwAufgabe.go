package dhbw

// Foo erwartet einen String und hängt diesen so oft, wie das Wort beim ersten Aufruf lang ist in voller länge hintereinander
// an diesen an.
// Vorraussetzung: Löse Rekursiv
// Foo Liefert das Ergebnis.

func FooS(s string) string {
	n := len(s)
	var result string
	for i := 0; i < n; i++ {
		result = result + s
	}
	return result

}

func FooR(s string) string {
	return Foo(s, len(s))

}
func Foo(s string, n int) string {
	if n == 1 {
		return s
	}
	return s + Foo(s, n-1)
}
