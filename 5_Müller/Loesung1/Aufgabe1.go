package Loesung1

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion Mult.
PUNKTE: 5
BEWERTUNG:
Rekursiv!!!!
*/

// MultFive erwartet eine Liste von float64 und eine Liste von int und liefert
// das Produkt der Elemente an den selben Stellen.
// Die Aufgabe muss rekursiv gelöst werden. Falls eine Stelle nicht
// existiert soll mit 0 multipliziert werden.
func Mult(l1 []float64, l2 []int) []float64 {
	if len(l1) == 0 || len(l2) == 0 {
		return []float64{}
	}
	return append([]float64{l1[0] * float64(l2[0])}, Mult(l1[1:], l2[1:])...)
}
