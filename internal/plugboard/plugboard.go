package plugboard

import (
	"errors"
)

var (
	ErrInvalidConnection = errors.New("invalid connection")
	ErrDuplicatePlug     = errors.New("duplicate plug")
)

type Connection struct {
	A rune `json:"a"`
	B rune `json:"b"`
}

type Plugboard struct {
	w map[int]int
}

func basic() map[int]int {
	m := make(map[int]int)
	for i := 0; i < 26; i++ {
		m[i] = i
	}
	return m
}

func NewPlugboard(conns []*Connection) (*Plugboard, error) {
	if len(conns) < 1 {
		return &Plugboard{w: basic()}, nil
	}

	var (
		wires = basic()
		set   = map[int]struct{}{}
	)
	for _, c := range conns {
		if c.A == 0 || c.B == 0 {
			return nil, ErrInvalidConnection
		}

		var (
			a = int(c.A) - 65
			b = int(c.B) - 65
		)

		if _, ok := set[a]; ok {
			return nil, ErrDuplicatePlug
		}
		set[a] = struct{}{} // add since it doesn't exist

		if _, ok := set[b]; ok {
			return nil, ErrDuplicatePlug
		}
		set[b] = struct{}{} // add since it doesn't exist

		wires[a] = b
		wires[b] = a
	}

	p := &Plugboard{w: wires}
	return p, nil
}

func (pb *Plugboard) Forward(c int) int {
	i, ok := pb.w[c]
	if !ok {
		return c
	}

	return i
}
