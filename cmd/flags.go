package cmd

import (
	"time"

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
		Name:        "ami-type",
		Aliases:     []string{"t"},
		DefaultText: "\"AL2023_x86_64_STANDARD\" or \"AUTO_MODE_STANDARD_x86_64\"",
		Usage:       "AMI Type for the AMI",
	},
	&cli.StringFlag{
		Name:    "kubernetes-version",
		Aliases: []string{"V"},
		Value:   "1.34",
		Usage:   "Kubernetes version for AMI",
	},
	&cli.StringFlag{
		Name:    "release-date",
		Aliases: []string{"d"},
		Value:   "",
		Usage:   "Release date with [yyyy], [yyyymm] or [yyyymmdd] format",
	},
	&cli.DurationFlag{
		Name:  "timeout",
		Value: 30 * time.Second,
		Usage: "Request timeout duration",
	},
	&cli.BoolFlag{
		Name:  "auto-mode",
		Value: false,
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
