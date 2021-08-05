package enigma

import (
	"strconv"
	"strings"

	"github.com/elliottpolk/enigma/internal/machine"
	"github.com/elliottpolk/enigma/internal/plugboard"
	"github.com/elliottpolk/enigma/internal/reflector"
	"github.com/elliottpolk/enigma/internal/rotor"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type config struct {
	Rotors map[string]*struct {
		Name string `json:"name"`
		Pos  int    `json:"pos"`
		Ring int    `json:"ring"`
	} `json:"rotors"`

	Reflector string   `json:"reflector"`
	Plugboard []string `json:"plugboard"`
}

type Enigma struct {
	*machine.Machine
}

//func NewMachine(c, s []byte, sf string) (*Enigma, error) {
func NewMachine(c []byte) (*Enigma, error) {
	var in map[string]*config
	if err := yaml.Unmarshal(c, &in); err != nil {
		return nil, errors.Wrap(err, "unable to parse config file")
	}

	var (
		cfg = in["machine"]

		lr = cfg.Rotors["left"]
		mr = cfg.Rotors["middle"]
		rr = cfg.Rotors["right"]

		conns = make([]*plugboard.Connection, 0)
	)

	for _, conn := range cfg.Plugboard {
		if len(conn) < 2 {
			return nil, errors.Errorf("invalid plugboard config %s", conn)
		}

		_conn := []rune(strings.ToUpper(conn))

		var (
			a = _conn[0]
			b = _conn[1]
		)

		// confirm 'a' is not a number
		if _, err := strconv.Atoi(string(a)); err == nil {
			return nil, errors.Errorf("invalid plugboard config %s", conn)
		}

		// confirm 'b' is not a number
		if _, err := strconv.Atoi(string(b)); err == nil {
			return nil, errors.Errorf("invalid plugboard config %s", conn)
		}

		conns = append(conns, &plugboard.Connection{a, b})
	}

	pb, err := plugboard.NewPlugboard(conns)
	if err != nil {
		return nil, err
	}

	return &Enigma{
		&machine.Machine{
			Left:      rotor.Create(lr.Name, lr.Pos, lr.Ring),
			Middle:    rotor.Create(mr.Name, mr.Pos, mr.Ring),
			Right:     rotor.Create(rr.Name, rr.Pos, rr.Ring),
			Reflector: reflector.Create(cfg.Reflector),
			Plugboard: pb,
		},
	}, nil
}
