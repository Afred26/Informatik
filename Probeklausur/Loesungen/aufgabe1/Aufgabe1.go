package aufgabe1

// Du erhällst zwei Startwerte einer Zahlenfolge x & y
// mit der Gleichung x + y = z1 und y + z1 = z2 wird die Zahlenfolge fortgesetzt
// berechen den n-ten term der Folge

func Foo(x, y, n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return x
	}
	if n == 2 {
		return y
	}

	return Foo(y, x+y, n-1)
}
