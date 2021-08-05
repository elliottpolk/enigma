package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/elliottpolk/enigma"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	cli "github.com/urfave/cli/v2"
	altsrc "github.com/urfave/cli/v2/altsrc"
)

var (
	version  string
	compiled string = fmt.Sprint(time.Now().Unix())
	githash  string

	cfgFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "config",
		Aliases: []string{"confg", "cfg", "c"},
		Usage:   "path to config file",
	})
)

func main() {
	ct, err := strconv.ParseInt(compiled, 0, 0)
	if err != nil {
		panic(err)
	}

	app := cli.App{
		Name:      "enigma",
		Copyright: "Copyright © 2021 Elliott Polk",
		Version:   fmt.Sprintf("%s | compiled %s | commit %s", version, time.Unix(ct, -1).Format(time.RFC3339), githash),
		Compiled:  time.Unix(ct, -1),
		Flags: []cli.Flag{
			cfgFlag,
		},
		Action: func(ctx *cli.Context) error {
			cfg, err := ioutil.ReadFile(ctx.String(cfgFlag.Name))
			if err != nil {
				return cli.Exit(errors.Wrap(err, "unable to read in config file"), 1)
			}

			if len(cfg) < 1 {
				return cli.Exit("a valid config file must be provided", 1)
			}

			m, err := enigma.NewMachine(cfg)
			if err != nil {
				return cli.Exit(errors.Wrap(err, "unable to create new machine"), 1)
			}

			r := bufio.NewReader(os.Stdin)
			for {
				fmt.Print("> ")

				in, err := r.ReadString('\n')
				if err != nil {
					log.Error(err)
					continue
				}

				// remove all lead & trailing spaces and force to uppercase
				in = strings.ToUpper(strings.TrimSpace(in))

				fmt.Printf("> %s\n\n", m.Encrypt(in))
			}

			return nil
		},
	}

	app.Run(os.Args)
}
