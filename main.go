package main

import (
	"os"

	"github.com/guessi/eks-ami-finder/cmd"
	"github.com/guessi/eks-ami-finder/pkg/constants"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:    constants.NAME,
		Usage:   constants.USAGE,
		Version: constants.VERSION,
		Flags:   cmd.Flags,
		Action: func(c *cli.Context) error {
			cmd.Wrapper(c)
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		os.Exit(1)
	}
}
