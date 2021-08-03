package wiring

type Wiring []int

func Decode(value []rune) Wiring {
	wiring := make(Wiring, len(value))
	for i, c := range value {
		wiring[i] = (int(c) - 65)
	}

	return wiring
}

func Inverse(value []int) Wiring {
	inverse := make(Wiring, len(value))
	for i, c := range value {
		inverse[c] = i
	}
	return inverse
}
