package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/flaflasun/wcnt/cmd"
)

func main() {
	app := cli.NewApp()
	app.Name = "wcnt"
	app.Version = Version
	app.Usage = "print the word count"
	app.Author = "flaflasun"

	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "reverse, r",
			Usage: "switch to ascending order",
		},
		cli.BoolFlag{
			Name:  "ignorecase, i",
			Usage: "the case of normal letters is ignored",
		},
		cli.StringFlag{
			Name:  "ignorechar, c",
			Value: ",.:;()!?",
			Usage: "the characters are ignored",
		},
		cli.StringFlag{
			Name:  "delimiter, d",
			Value: "",
			Usage: "additional delimiter",
		},
	}

	app.Action = cmd.Wcnt
	app.Run(os.Args)
}
