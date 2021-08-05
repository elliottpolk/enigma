package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
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

	stateFlag = altsrc.NewStringFlag(&cli.StringFlag{
		Name:    "state",
		Aliases: []string{"s"},
		Usage:   "path to read/write a state file",
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
			stateFlag,
		},
		Action: func(ctx *cli.Context) error {
			cfg, err := ioutil.ReadFile(ctx.String(cfgFlag.Name))
			if err != nil {
				return cli.Exit(errors.Wrap(err, "unable to read in config file"), 1)
			}

			if len(cfg) < 1 {
				return cli.Exit("a valid config file must be provided", 1)
			}

			sf := ctx.String(stateFlag.Name)
			if len(sf) < 1 {
				return cli.Exit("a valid state file must be provided", 1)
			}

			state, err := ioutil.ReadFile(sf)
			if err != nil {
				return cli.Exit(errors.Wrap(err, "unable to read in state file"), 1)
			}

			// capture kill signal to write the state
			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			go func() {
				for sig := range c {
					_ = sig
					fmt.Println()
					log.Info("capturing and writing current state...")

					// TODO:
					// - capture state
					// - write to state file

					log.Info("state written")
					os.Exit(0)
				}
			}()

			m, err := enigma.NewMachine(cfg, state, sf)
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

				fmt.Println(m.Encrypt(in))
			}

			return nil
		},
	}

	app.Run(os.Args)
}
