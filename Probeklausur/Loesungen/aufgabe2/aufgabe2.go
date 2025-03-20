package aufgabe2

// Die Summe der ersten 10 Quadartzahlen ist:
// 1^2 + 2^2 + ... + 10^2 = 385.
// Das Quadart der Summe der ersten 10 Zahlen ist:
// (1 + 2 + ... + 10)^2 = 55^2 = 3025.
// Die Diverenz der beiden ist: 3025 - 385 = 2640.
// Finde die Diverenz der ersten n Zahlen

func Difference(n int) int {
	return Pow2(Sumdown(n)) - SumSquare(n)
}

func SumSquare(n int) int {
	if n <= 0 {
		return 0
	}
	return SumSquare(n-1) + Pow2(n)
}

func Sumdown(n int) int {
	if n <= 0 {
		return 0
	}
	return Sumdown(n-1) + n
}

func Pow2(n int) int {
	return n * n
}
