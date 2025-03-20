package aufgabe3

// Wenn 2 die erste Primzahl ist, 3 die zweite, 5 die dritte
// gib die n-te Primzahl an
// n ist kleiner als 10000

func Prime(n int) int {
	Prims := Primsbelow(200000)
	return Prims[n-1]
}

func Primsbelow(max int) []int {
	prims := []int{}
	for i := 2; i < max; i++ {
		if !Div(prims, i) {
			prims = append(prims, i)
		}
	}

	return prims
}

func Div(list []int, d int) bool {
	for _, el := range list {
		if d%el == 0 {
			return true
		}
	}
	return false
}
