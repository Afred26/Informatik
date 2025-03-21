package aufgabe2

// EverythingBevorSubstring nimmt eine Liste von Strings und ein Substring als Eingabe.
// Es gibt einen neue Liste zurück, mit Teilstrings bis zu dem Substring und/oder Strings ohne Substring.
// Sollte die Liste leer sein gebe eine Liste zurück die nur den Substring enthaelt.
func EverythingBevorSubstring(list []string, sub string) []string {

	return []string{}
}

// Index gibt den Index des Substrings in einem String zurück.
func index(str, sub string) int {
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
