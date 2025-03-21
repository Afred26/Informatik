package Aufgabe3

import "math"

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
	Grenze := int(math.Sqrt(float64(x)))
	for i := 2; i <= Grenze; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}
