package aufgabe2

import "strings"

// EverythingBevorSubstring nimmt eine Liste von Strings und ein Substring als Eingabe.
// Es gibt einen neue Liste zurück, mit Teilstrings bis zu dem Substring und/oder Strings ohne Substring.
// Sollte die Liste leer sein gebe eine Liste zurück die nur den Substring enthaelt.
func EverythingBevorSubstring(list []string, sub string) []string {
	result := []string{}
	if len(list) == 0 {
		return []string{sub}
	}
	for _, el := range list {
		pos := strings.Index(el, sub)
		if pos == -1 {
			result = append(result, el)
		} else if pos > 0 {
			result = append(result, el[:pos])
		}

	}
	return result
}
