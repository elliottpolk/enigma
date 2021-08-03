package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

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
		Usage:   "optional path to config file",
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

			cfg, err := getCfg(ctx)
			if err != nil {
				return cli.Exit(err, 1)
			}

			var (
				lr = cfg.Rotors["left"]
				mr = cfg.Rotors["middle"]
				rr = cfg.Rotors["right"]
			)

			log.Info("not done yet")

			return nil
		},
	}

	app.Run(os.Args)
}
