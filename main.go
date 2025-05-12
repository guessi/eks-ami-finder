package main

import (
	"context"
	"fmt"
	"os"
	"regexp"

	"github.com/guessi/eks-ami-finder/cmd"
	"github.com/guessi/eks-ami-finder/pkg/constants"
	"github.com/urfave/cli/v3"
)

func showVersion() {
	r, _ := regexp.Compile(`v[0-9]\.[0-9]+\.[0-9]+`)
	versionInfo := r.FindString(constants.GitVersion)
	fmt.Println("eks-ami-finder", versionInfo)
	fmt.Println(" Git Commit:", constants.GitVersion)
	fmt.Println(" Build with:", constants.GoVersion)
	fmt.Println(" Build time:", constants.BuildTime)
}

func main() {
	app := &cli.Command{
		Name:    constants.NAME,
		Usage:   constants.USAGE,
		Version: constants.GitVersion,
		Flags:   cmd.Flags,
		Action: func(ctx context.Context, c *cli.Command) error {
			cmd.Wrapper(c)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "version",
				Aliases: []string{"v"},
				Usage:   "Print version number",
				Action: func(context.Context, *cli.Command) error {
					showVersion()
					return nil
				},
			},
		},
	}

	if err := app.Run(context.Background(), os.Args); err != nil {
		os.Exit(1)
	}
}
