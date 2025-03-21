package aufgabe2

// EverythingBevorSubstring nimmt eine Liste von Strings und ein Substring als Eingabe.
// Es gibt eine neue Liste zurück, mit Teilstrings bis zu dem Substring und/oder Strings ohne Substring.
// Sollte die Liste leer sein, gebe eine Liste zurück, die nur den Substring enthält.
func EverythingBevorSubstring_loesung(list []string, sub string) []string {
	if len(list) == 0 {
		return []string{sub}
	}

	var result []string
	for _, str := range list {
		if idx := indexof(str, sub); idx != -1 {
			if idx > 0 { // Skip appending empty strings
				result = append(result, str[:idx])
			}
		} else {
			result = append(result, str)
		}

	}

	return result
}

// Index gibt den Index des Substrings in einem String zurück.
func indexof(str, sub string) int {
	for i := 0; i < len(str); i++ {
		if len(str)-i < len(sub) {
			return -1
		}

		if str[i:i+len(sub)] == sub {
			return i
		}
	}

	return -1
}
