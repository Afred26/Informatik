package Loesung3

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion IsPrim.
PUNKTE: 5
BEWERTUNG:
*/

// IsPrim erwartet eine Zahl und liefert ob die Zahl
//  eine Primzahl ist.

func IsPrim(x int) bool {
	if x < 2 {
		return false
	}
	for i := x - 1; i > 1; i-- {
		if x%i == 0 {
			return false
		}
	}
	return true
}
