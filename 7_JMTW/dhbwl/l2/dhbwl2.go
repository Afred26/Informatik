package dhbwl

// Foo erwartet einen String und hängt diesen so oft, wie das Wort beim ersten Aufruf lang ist in voller länge hintereinander
// an diesen an.
// Vorraussetzung: Löse Rekursiv
// Foo Liefert das Ergebnis.

func Foo(s string, c int) string { // Weiterer Übergabe Parameter Int zur Längen Definition
	if c == 0 { // Rekursionsanker
		return ""
	}
	return s + Foo(s, c-1) // Rekursionsschritt

}
