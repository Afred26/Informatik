package dhbwl

// Foo erwartet einen String und hängt diesen so oft, wie das Wort beim ersten Aufruf lang ist in voller länge hintereinander
// an diesen an.
// Vorraussetzung: Löse Rekursiv
// Foo Liefert das Ergebnis.

func Foo(s, original string) string { // Weiterer Übergabe Parameter zur Längen Definition
	if len(original) == 0 { // Rekursions Anker
		return ""
	}
	return s + Foo(s, original[1:]) // Rekursionsschritt

}
