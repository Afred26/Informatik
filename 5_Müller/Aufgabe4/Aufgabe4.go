package Aufgabe4

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
	Numbers := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for _, el := range Numbers {
		if el == s[:1] {
			return FilterNumbers(s[1:])
		}
	}
	return s[:1] + FilterNumbers(s[1:])
}
