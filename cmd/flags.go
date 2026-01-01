package cmd

import (
	"context"
	"fmt"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/guessi/eks-ami-finder/pkg/constants"
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
		Action: func(ctx context.Context, c *cli.Command, v string) error {
			if v == "" {
				return nil // Empty is allowed (will be auto-resolved)
			}

			if len(v) != 12 {
				return fmt.Errorf("owner-id must be a 12-digit AWS account ID")
			}

			// Check if all characters are digits
			for _, char := range v {
				if char < '0' || char > '9' {
					return fmt.Errorf("owner-id must be a 12-digit AWS account ID")
				}
			}

			return nil
		},
	},
	&cli.StringFlag{
		Name:        "ami-type",
		Aliases:     []string{"t"},
		DefaultText: "\"AL2023_x86_64_STANDARD\" or \"AUTO_MODE_STANDARD_x86_64\"",
		Usage:       "AMI Type for the AMI",
		Action: func(ctx context.Context, c *cli.Command, v string) error {
			if v == "" {
				return nil // Empty is allowed (will be auto-resolved based on auto-mode)
			}

			// Check context-aware validation based on auto-mode flag
			if c.Bool("auto-mode") {
				if !slices.Contains(constants.ValidAmiTypes["AUTO_MODE"], v) {
					return fmt.Errorf("invalid ami-type '%s' for auto-mode. Valid types: %s", v, strings.Join(constants.ValidAmiTypes["AUTO_MODE"], ", "))
				}
			} else {
				if !slices.Contains(constants.ValidAmiTypes["DEFAULT"], v) {
					return fmt.Errorf("invalid ami-type '%s'. Supported ami-type could be found at https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html", v)
				}
			}

			return nil
		},
	},
	&cli.StringFlag{
		Name:    "kubernetes-version",
		Aliases: []string{"V"},
		Value:   "1.34",
		Usage:   "Kubernetes version for AMI",
		Action: func(ctx context.Context, c *cli.Command, v string) error {
			parts := strings.Split(v, ".")
			if len(parts) != 2 {
				return fmt.Errorf("invalid Kubernetes version format. Expected format: X.Y (e.g., 1.33)")
			}

			major, err1 := strconv.Atoi(parts[0])
			minor, err2 := strconv.Atoi(parts[1])
			if err1 != nil || err2 != nil {
				return fmt.Errorf("invalid Kubernetes version format. Expected format: X.Y (e.g., 1.33)")
			}

			// The first Amazon EKS version was 1.10
			// - https://aws.amazon.com/blogs/aws/amazon-eks-now-generally-available/
			if major != 1 || minor < 10 {
				return fmt.Errorf("the very first Amazon EKS version was 1.10, so there would have no Amazon EKS %s", v)
			}

			return nil
		},
	},
	&cli.StringFlag{
		Name:    "release-date",
		Aliases: []string{"d"},
		Value:   "",
		Usage:   "Release date with [yyyy], [yyyymm] or [yyyymmdd] format",
		Action: func(ctx context.Context, c *cli.Command, v string) error {
			if v == "" {
				return nil // Empty is allowed
			}

			// Validate format: yyyy, yyyymm, or yyyymmdd
			if len(v) != 4 && len(v) != 6 && len(v) != 8 {
				return fmt.Errorf("invalid release-date format. Expected [yyyy], [yyyymm] or [yyyymmdd]")
			}

			// Check if all characters are digits
			for _, char := range v {
				if char < '0' || char > '9' {
					return fmt.Errorf("invalid release-date format. Expected [yyyy], [yyyymm] or [yyyymmdd]")
				}
			}

			return nil
		},
	},
	&cli.DurationFlag{
		Name:  "timeout",
		Value: 30 * time.Second,
		Usage: "Request timeout duration",
		Action: func(ctx context.Context, c *cli.Command, v time.Duration) error {
			if v <= 0 {
				return fmt.Errorf("timeout must be greater than 0")
			}
			return nil
		},
	},
	&cli.BoolFlag{
		Name:  "auto-mode",
		Value: false,
	},
	&cli.BoolFlag{
		Name:  "include-deprecated",
		Value: false,
	},
	&cli.IntFlag{
		Name:    "max-results",
		Aliases: []string{"n"},
		Value:   20,
		Action: func(ctx context.Context, c *cli.Command, v int) error {
			if v <= 0 {
				return fmt.Errorf("max-results must be greater than 0")
			}
			return nil
		},
	},
	&cli.BoolFlag{
		Name:  "debug",
		Value: false,
	},
}
