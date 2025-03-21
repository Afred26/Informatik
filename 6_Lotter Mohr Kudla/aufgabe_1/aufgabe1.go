package aufgabe1

// StringsWithoutDigit erwartet eine Liste von Strings und liefert
// eine Liste an Strings welche keine Zahlen beinhalten.
// Sollten keine Strings ohne Zahlen gegebensein soll er eine leere Liste zurückgeben.
func StringsWithoutDigit(list []string) []string {

	result := []string{}
	for _, el := range list {
		if WithoutDigit(el) {
			result = append(result, el)
		}
	}
	return result
}

func WithoutDigit(s string) bool {
	if s == "" {
		return true
	}
	if s[0] >= '0' && s[0] <= '9' {
		return false
	}
	return WithoutDigit(s[1:])
}
