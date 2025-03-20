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
	if Pow2(p.a)+Pow2(p.b) == Pow2(p.c) {
		if p.a < p.b && p.b < p.c {
			return true
		}
	}
	return false

}

func (p Pythagoreantriplet) SumThousen() bool {
	if p.a+p.b+p.c == 1000 {
		return true
	} else {
		return false
	}
}

func FindTriplets(n int) Pythagoreantriplet {
	list := GetIntList(n)
	p := Pythagoreantriplet{a: 0, b: 0, c: 0}
	result := Pythagoreantriplet{}

	for _, ela := range list {
		p.a = ela
		for _, elb := range list {
			p.b = elb
			if p.a+p.b < n {
				p.c = n - (p.a + p.b)
				if p.IsPythagorean() {
					result = p
				}
			}
		}
	}

	return result

}
func GetIntList(n int) []int {
	if n < 0 {
		return []int{}
	}
	return append(GetIntList(n-1), n)
}

func Pow2(n int) int {
	return n * n
}
