package machine

import (
	"github.com/elliottpolk/enigma/internal/plugboard"
	"github.com/elliottpolk/enigma/internal/reflector"
	"github.com/elliottpolk/enigma/internal/rotor"
)

type Machine struct {
	left   *rotor.Rotor
	middle *rotor.Rotor
	right  *rotor.Rotor

	r *reflector.Reflector
	p *plugboard.Plugboard
}

func (m *Machine) rotate() {
	if m.middle.AtNotch() {
		m.middle.Turnover()
		m.left.Turnover()
	} else if m.right.AtNotch() {
		m.middle.Turnover()
	}

	m.right.Turnover()
}

func (m *Machine) encrypt(c int) int {
	m.rotate()

	// plugboard in
	c = m.p.Forward(c)

	// right to left
	c = m.right.Forward(c)
	c = m.middle.Forward(c)
	c = m.left.Forward(c)

	// reflector
	c = m.r.Forward(c)

	// left to right
	c = m.left.Backward(c)
	c = m.middle.Backward(c)
	c = m.right.Backward(c)

	// plugboard out
	return m.p.Forward(c)
}

func (m *Machine) Encrypt(input string) string {

	in := []rune(input)
	buf := make([]rune, len(in))

	for i, c := range in {
		buf[i] = rune(m.encrypt(int(c)-65) + 65)
	}

	return string(buf)
}
