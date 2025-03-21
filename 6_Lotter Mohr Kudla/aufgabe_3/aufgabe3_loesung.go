package aufgabe3

//AddEachElement erwartetzwei Listen von Integern und gibt
//eine Liste zurück, die die Summe der Elemente an der gleichen Position enthält.
//Sollte eine liste mehr Elemente enthalten als die andere,
//werden die restlichen zahlen mit den letzten uahl der kurzeren Liste addiert.
//Sollten beide Listen leer sein, dann  soll er die andere liste zurückgeben.

func AddEachElement_loesung(l1, l2 []int) []int {

	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}

	if len(l1) > len(l2) {
		if len(l2) == 1 {
			return append([]int{l1[0] + l2[0]}, AddEachElement(l1[1:], l2)...)
		}
	}

	if len(l1) < len(l2) {
		if len(l1) == 1 {
			return append([]int{l1[0] + l2[0]}, AddEachElement(l1, l2[1:])...)
		}

	}
	return append([]int{l1[0] + l2[0]}, AddEachElement(l1[1:], l2[1:])...)
}
