package aufgabe4

// Ein Pythagoreisches Trippel sind drei natürliche Zahlen, für die gilt  a < b < c & a^2 + b^2 = c^2.
// Z.B.: 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
// Prüfe ob drei gegeben Zahlen ein solches Trippel sind
// Und fide ein Trippel sodas a + b + c = 1000

type Pythagoreantriplet struct {
	a int
	b int
	c int
}

func (p Pythagoreantriplet) IsPythagorean() bool {
	return false

}

func FindTriplets(n int) Pythagoreantriplet {
	return Pythagoreantriplet{a: 3, b: 4, c: 5}
}
