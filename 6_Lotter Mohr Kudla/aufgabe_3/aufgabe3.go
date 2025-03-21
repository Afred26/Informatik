package aufgabe3

//AddEachElement erwartetzwei Listen von Integern und gibt
//eine Liste zurück, die die Summe der Elemente an der gleichen Position enthält.
//Sollte eine liste mehr Elemente enthalten als die andere,
//werden die restlichen Zahlen mit den letzten Zahl der kurzeren Liste addiert.
//Sollten beide Listen leer sein, dann  soll er die andere liste zurückgeben.

func AddEachElement(l1, l2 []int) []int {

	min := Min(len(l1), len(l2))
	max := Max(len(l1), len(l2))
	result := make([]int, max)
	if len(l1) == 0 {
		return l2
	}
	if len(l2) == 0 {
		return l1
	}

	if len(l1) == len(l2) {
		for i := range result {
			result[i] = l1[i] + l2[i]
		}
	} else if len(l1) > len(l2) {
		for i := 0; i < min; i++ {
			result[i] = l1[i] + l2[i]
		}
		for i := min; i < max; i++ {
			result[i] = l1[i] + l2[min-1]
		}

	} else if len(l2) > len(l1) {
		for i := 0; i < min; i++ {
			result[i] = l1[i] + l2[i]
		}
		for i := min; i < max; i++ {
			result[i] = l1[min-1] + l2[i]
		}

	}
	return result
}

func Max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}

}

func Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}

}
