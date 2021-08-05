package rotor

import (
	"strings"

	"github.com/elliottpolk/enigma/internal/wiring"
)

const limit int = 26

type Rotor struct {
	Name  string      `json:"Name"`
	Wires *wiring.Set `json:"wires"`
	Pos   int         `json:"pos"`
	Notch int         `json:"notch"`
	Ring  int         `json:"ring"`
}

func Create(n string, p, r int) *Rotor {
	var (
		w string

		notch int
		name  string = strings.ToUpper(n)
	)

	switch name {
	case "I":
		w = "EKMFLGDQVZNTOWYHXUSPAIBRCJ"
		notch = 16

	case "II":
		w = "AJDKSIRUXBLHWTMCQGZNPYFVOE"
		notch = 4

	case "III":
		w = "BDFHJLCPRTXVZNYEIWGAKMUSQO"
		notch = 21

	case "IV":
		w = "ESOVPZJAYQUIRHXLNFTGKDCMWB"
		notch = 9

	case "V":
		w = "VZBRGITYUPSDNHLXAWMJQOFECK"
		notch = 25

	case "VI":
		w = "JPGVOUMFYQBENHZRDKASXLICTW"
		notch = 0

	case "VII":
		w = "NZJHGRCXMYSWBOUFAIVLPEKQDT"
		notch = 0

	case "VIII":
		w = "FKQHTLXOCBJSPDZRAMEWNIUYGV"
		notch = 0

	default:
		fwd := wiring.Decode([]rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ"))
		return &Rotor{
			Name: "Identity",
			Wires: &wiring.Set{
				Forward:  fwd,
				Backward: wiring.Inverse(fwd),
			},
			Pos:   p,
			Notch: 0,
			Ring:  r,
		}
	}

	fwd := wiring.Decode([]rune(w))
	return &Rotor{
		Name: name,
		Wires: &wiring.Set{
			Forward:  fwd,
			Backward: wiring.Inverse(fwd),
		},
		Pos:   p,
		Notch: notch,
		Ring:  r,
	}
}

func encipher(k, pos, ring int, mapping []int) int {
	shift := pos - ring
	return (mapping[(k+shift+limit)%limit] - shift + limit) % limit
}

func (r *Rotor) Forward(c int) int {
	return encipher(c, r.Pos, r.Ring, r.Wires.Forward)
}

func (r *Rotor) Backward(c int) int {
	return encipher(c, r.Pos, r.Ring, r.Wires.Backward)
}

func (r *Rotor) AtNotch() bool {
	switch r.Name {
	case "VI", "VII", "VIII":
		return r.Pos == 12 || r.Pos == 25

	default:
		return r.Notch == r.Pos
	}
}

func (r *Rotor) Turnover() {
	r.Pos = (r.Pos + 1) % limit
}
