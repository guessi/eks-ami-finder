package cmd

import (
	"github.com/urfave/cli/v2"
)

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:     "region",
		Aliases:  []string{"r"},
		Value:    "us-east-1",
		Usage:    "Region for the AMI",
		Required: false,
	},
	&cli.StringFlag{
		Name:     "owner-id",
		Aliases:  []string{"o"},
		Value:    "",
		Usage:    "Owner ID of the AMI",
		Required: false,
	},
	&cli.StringFlag{
		Name:     "ami-type",
		Aliases:  []string{"t"},
		Value:    "x86_64",
		Usage:    "x86_64, x86_64-gpu, arm64",
		Required: false,
	},
	&cli.StringFlag{
		Name:     "ami-family",
		Aliases:  []string{"f"},
		Value:    "AL2",
		Usage:    "AL2, AL2023",
		Required: false,
	},
	&cli.StringFlag{
		Name:     "kubernetes-version",
		Aliases:  []string{"V"},
		Value:    "1.30",
		Usage:    "Kubernetes version for AMI",
		Required: false,
	},
	&cli.StringFlag{
		Name:     "release-date",
		Aliases:  []string{"d"},
		Value:    "",
		Usage:    "Release date with [yyyy], [yyyymm] or [yyyymmdd] format",
		Required: false,
	},
	&cli.BoolFlag{
		Name:     "include-deprecated",
		Value:    false,
		Required: false,
	},
	&cli.StringFlag{
		Name:     "max-results",
		Aliases:  []string{"n"},
		Value:    "20",
		Required: false,
	},
	&cli.BoolFlag{
		Name:     "debug",
		Value:    false,
		Required: false,
	},
}
