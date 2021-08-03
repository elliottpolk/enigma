package reflector

import "github.com/elliottpolk/enigma/wiring"

type Reflector struct {
	forward wiring.Wiring
}

func (r *Reflector) Forward(c int) int {
	return r.forward[c]
}

func Create(n string) *Reflector {
	var w string

	switch n {
	case "B":
		w = "YRUHQSLDPXNGOKMIEBFZCWVJAT"

	case "C":
		w = "FVPJIAOYEDRZXWGCTKUQSBNMHL"

	default:
		w = "ZYXWVUTSRQPONMLKJIHGFEDCBA"
	}

	return &Reflector{forward: wiring.Decode([]rune(w))}
}
