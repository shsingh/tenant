package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

func noArgs(c *cli.Context) {
	cli.ShowAppHelp(c)
}

func main() {
	app := &cli.App{
		Name:    "tenant",
		Version: "0.0.1",
		Usage:   "Create a virtual shell with environment variables.",
		Authors: []*cli.Author{
			&cli.Author{
				Name:  "Shain Singh",
				Email: "shsingh@ieee.org",
			},
		},
		Action: func(c *cli.Context) error {
			noArgs(c)
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
