package rotor

import (
	"strings"

	"github.com/elliottpolk/enigma/wiring"
)

const limit int = 26

type Rotor struct {
	Name string

	forward  wiring.Wiring
	backward wiring.Wiring

	rpos int
	npos int

	ring int
}

func Create(n string, rp, rs int) *Rotor {

	var (
		w string

		npos int
		name string = strings.ToUpper(n)
	)

	switch name {
	case "I":
		w = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
		npos = 16

	case "II":
		w = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
		npos = 4

	case "III":
		w = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
		npos = 21

	case "IV":
		w = "ESOVPZJAYQUIRHXLNFTGKDCMWB"
		npos = 9

	case "V":
		w = "VZBRGITYUPSDNHLXAWMJQOFECK"
		npos = 25

	case "VI":
		w = "JPGVOUMFYQBENHZRDKASXLICTW"
		npos = 0

	case "VII":
		w = "NZJHGRCXMYSWBOUFAIVLPEKQDT"
		npos = 0

	case "VIII":
		w = "FKQHTLXOCBJSPDZRAMEWNIUYGV"
		npos = 0

	default:
		fwd := wiring.Decode([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
		return &Rotor{
			Name:     "Identity",
			forward:  fwd,
			backward: wiring.Inverse(fwd),
			rpos:     rp,
			npos:     0,
			ring:     rs,
		}
	}

	fwd := wiring.Decode([]rune(w))
	return &Rotor{
		Name:     name,
		forward:  fwd,
		backward: wiring.Inverse(fwd),
		rpos:     rp,
		npos:     npos,
		ring:     rs,
	}
}

func encipher(k, pos, ring int, mapping []int) int {
	shift := pos - ring
	return (mapping[(k+shift+limit)%limit] - shift + limit) % limit
}

func (r *Rotor) Forward(c int) int {
	return encipher(c, r.rpos, r.ring, r.forward)
}

func (r *Rotor) Backward(c int) int {
	return encipher(c, r.rpos, r.ring, r.backward)
}

func (r *Rotor) AtNotch() bool {
	switch r.Name {
	case "VI", "VII", "VIII":
		return r.rpos == 12 || r.rpos == 25

	default:
		return r.npos == r.rpos
	}
}

func (r *Rotor) Turnover() {
	r.rpos = (r.rpos + 1) % limit
}
