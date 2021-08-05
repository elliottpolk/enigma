package machine

import (
	"github.com/elliottpolk/enigma/internal/plugboard"
	"github.com/elliottpolk/enigma/internal/reflector"
	"github.com/elliottpolk/enigma/internal/rotor"
)

type Machine struct {
	Left   *rotor.Rotor
	Middle *rotor.Rotor
	Right  *rotor.Rotor

	Reflector *reflector.Reflector
	Plugboard *plugboard.Plugboard
}

func (m *Machine) rotate() {
	if m.Middle.AtNotch() {
		m.Middle.Turnover()
		m.Left.Turnover()
	} else if m.Right.AtNotch() {
		m.Middle.Turnover()
	}

	m.Right.Turnover()
}

func (m *Machine) encrypt(c int) int {
	m.rotate()

	// plugboard in
	c = m.Plugboard.Forward(c)

	// right to left
	c = m.Right.Forward(c)
	c = m.Middle.Forward(c)
	c = m.Left.Forward(c)

	// reflector
	c = m.Reflector.Forward(c)

	// left to right
	c = m.Left.Backward(c)
	c = m.Middle.Backward(c)
	c = m.Right.Backward(c)

	// plugboard out
	return m.Plugboard.Forward(c)
}

func (m *Machine) Encrypt(input string) string {
	var (
		in  = []rune(input)
		buf = make([]rune, len(in))
	)

	for i, c := range in {
		if int(c) == 32 {
			buf[i] = c
		} else {
			buf[i] = rune(m.encrypt(int(c)-65) + 65)
		}
	}

	return string(buf)
}
