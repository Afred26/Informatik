package Aufgabe2

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion CountVocals.
PUNKTE: 5
BEWERTUNG:
*/

// CountVocals erwartet einen string und liefert
// die Anzahl an Vokale in einer []int{} in folgender Form:
// [a e i o u]
func CountVocals(s string) []int {
	Count := make([]int, 5)
	Vocalk := []rune{'a', 'e', 'i', 'o', 'u'}
	Vocalg := []rune{'A', 'E', 'I', 'O', 'U'}
	for i, el := range Vocalk {
		for _, b := range s {
			if el == b {
				Count[i]++
			}

		}
	}
	for i, el := range Vocalg {
		for _, b := range s {
			if el == b {
				Count[i]++
			}

		}
	}
	return Count
}
