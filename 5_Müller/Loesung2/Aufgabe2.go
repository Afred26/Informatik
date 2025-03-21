package Loesung2

import "strings"

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountUmlauts.
PUNKTE: 5
BEWERTUNG:
*/

// CountUmlauts erwartet einen string und liefert
// die Anzahl an Umlauten in einer []int{} in folgender Form:
// [a e i o u]
func CountUmlauts(s string) []int {
	s1 := strings.ToLower(s)
	a := strings.Count(s1, "a")
	e := strings.Count(s1, "e")
	i := strings.Count(s1, "i")
	o := strings.Count(s1, "o")
	u := strings.Count(s1, "u")
	return []int{a, e, i, o, u}
}
