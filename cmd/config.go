package main

import (
	"bufio"
	"io"
	"io/ioutil"
	"math"
	"os"

	"github.com/pkg/errors"
	cli "github.com/urfave/cli/v2"
	yaml "gopkg.in/yaml.v2"
)

type rotor struct {
	Name  string `json:"name"`
	Pos   int    `json:"pos"`
	Notch int    `json:"notch"`
}

type configuration struct {
	Rotors    map[string]*rotor `json:"rotors"`
	Reflector string            `json:"reflector"`
	Plugboard []string          `json:"plugboard"`
}

var (
	ErrNoPipe       = errors.New("no piped input")
	ErrDataTooLarge = errors.New("data to large")

	MaxData = (int(math.Pow10(7)) * 3)
)

func pipe() (string, error) {
	fi, err := os.Stdin.Stat()
	if err != nil {
		return "", errors.Wrap(err, "unable to stat stdin")
	}

	if fi.Mode()&os.ModeCharDevice != 0 {
		return "", ErrNoPipe
	}

	buf, res := bufio.NewReader(os.Stdin), make([]byte, 0)
	for {
		in, _, err := buf.ReadLine()
		if err != nil && err == io.EOF {
			break
		}
		res = append(res, in...)
		res = append(res, '\n')

		if len(res) > MaxData {
			return "", ErrDataTooLarge
		}
	}

	return string(res), nil
}

func getCfg(ctx *cli.Context) (*configuration, error) {
	var data []byte

	if f := ctx.String(cfgFlag.Name); len(f) > 0 {
		d, err := ioutil.ReadFile(f)
		if err != nil {
			return nil, errors.Wrap(err, "unable to read in config file")
		}

		data = make([]byte, len(d))
		copy(data, d)
	}

	var cfg map[string]*configuration
	if err := yaml.Unmarshal(data, &cfg); err != nil {
		return nil, errors.Wrap(err, "unable to parse config data")
	}

	return cfg["machine"], nil
}
