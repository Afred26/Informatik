package Loesung4

/*
AUFGABENSTELLUNG: Vervollständigen Sie die Funktion FilterNumbers.
PUNKTE: 5
BEWERTUNG:
Rekursiv!!!
*/

// FilterNumbers erwartet einen String und filtert alle Zahlen aus ihnen.
// Die Aufgabe muss rekursiv gelöst werden.

func FilterNumbers(s string) string {
	if len(s) == 0 {
		return ""
	}
	if rune(s[0]) == '0' || rune(s[0]) == '1' || rune(s[0]) == '2' || rune(s[0]) == '3' || rune(s[0]) == '4' ||
		rune(s[0]) == '5' || rune(s[0]) == '6' || rune(s[0]) == '7' || rune(s[0]) == '8' || rune(s[0]) == '9' {
		return FilterNumbers(s[1:])
	}
	return string(s[0]) + FilterNumbers(s[1:])
}
