package aufgabe1

//Lösung

// StringsWithoutDigit erwartet eine Liste von Strings und liefert
// eine Liste an Strings welche keine Zahlen beinhalten.
// Sollten keine Strings ohne Zahlen gegebensein soll er eine leere Liste zurückgeben.
func StringsWithoutDigit_loesung(list []string) []string {

	result := []string{}

	for _, el := range list {
		digitSeen := false
		for _, c := range el {
			if c == '0' || c == '1' || c == '2' || c == '3' || c == '4' || c == '5' || c == '6' || c == '7' || c == '8' || c == '9' {
				digitSeen = true
			}
		}
		if !digitSeen {
			result = append(result, el)
		}
	}
	return result
}
