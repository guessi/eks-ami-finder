package cmd

import (
	"github.com/urfave/cli/v3"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:    "region",
		Aliases: []string{"r"},
		Value:   "us-east-1",
		Usage:   "Region for the AMI",
	},
	&cli.StringFlag{
		Name:    "owner-id",
		Aliases: []string{"o"},
		Value:   "",
		Usage:   "Owner ID of the AMI",
	},
	&cli.StringFlag{
		Name:    "ami-type",
		Aliases: []string{"t"},
		Value:   "AL2023_x86_64_STANDARD",
		Usage:   "AMI Type for the AMI",
	},
	&cli.StringFlag{
		Name:    "kubernetes-version",
		Aliases: []string{"V"},
		Value:   "1.33",
		Usage:   "Kubernetes version for AMI",
	},
	&cli.StringFlag{
		Name:    "release-date",
		Aliases: []string{"d"},
		Value:   "",
		Usage:   "Release date with [yyyy], [yyyymm] or [yyyymmdd] format",
	},
	&cli.BoolFlag{
		Name:  "include-deprecated",
		Value: false,
	},
	&cli.StringFlag{
		Name:    "max-results",
		Aliases: []string{"n"},
		Value:   "20",
	},
	&cli.BoolFlag{
		Name:  "debug",
		Value: false,
	},
}
