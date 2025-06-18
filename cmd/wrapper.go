package cmd

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/guessi/eks-ami-finder/pkg/constants"
	"github.com/urfave/cli/v3"
)

func amiSearchInput(c *cli.Command) amiSearchInputSpec {
	maxResults, err := strconv.Atoi(c.String("max-results"))
	if err != nil {
		fmt.Printf("Invalid --max-results value: %s. Must be a valid integer.\n\n", c.String("max-results"))
		os.Exit(1)
	}

	return amiSearchInputSpec{
		AWS_REGION:         c.String("region"),
		AMI_OWNER_ID:       c.String("owner-id"),
		AMI_TYPE:           c.String("ami-type"),
		KUBERNETES_VERSION: c.String("kubernetes-version"),
		RELEASE_DATE:       c.String("release-date"),
		MAX_RESULTS:        maxResults,
		AUTO_MODE:          c.Bool("auto-mode"),
		INCLUDE_DEPRECATED: c.Bool("include-deprecated"),
		DEBUG_MODE:         c.Bool("debug"),
	}
}

func Wrapper(ctx context.Context, c *cli.Command) {
	ctx, cancel := context.WithTimeout(ctx, c.Duration("timeout"))
	defer cancel()

	r := amiSearchInput(c)

	// Set default AMI_TYPE based on AUTO_MODE
	if r.AMI_TYPE == "" {
		if r.AUTO_MODE {
			r.AMI_TYPE = "AUTO_MODE_STANDARD_x86_64"
		} else {
			r.AMI_TYPE = "AL2023_x86_64_STANDARD"
		}
	}

	// If region is specified but owner ID is missing or invalid, assume it is looking for EKS official image build
	if len(r.AWS_REGION) > 0 && (len(r.AMI_OWNER_ID) == 0 || len(r.AMI_OWNER_ID) != 12) {
		var mappings map[string]string
		switch {
		case strings.HasPrefix(r.AMI_TYPE, "AL2_"), strings.HasPrefix(r.AMI_TYPE, "AL2023_"):
			r.AMI_OWNER_ID = constants.AwsAccountMappingsAL["*"]
			mappings = constants.AwsAccountMappingsAL
		case strings.HasPrefix(r.AMI_TYPE, "BOTTLEROCKET_"):
			mappings = constants.AwsAccountMappingsBottlerocket
		case strings.HasPrefix(r.AMI_TYPE, "WINDOWS_"):
			mappings = constants.AwsAccountMappingsWindows
		}
		if mappings != nil {
			if v, ok := mappings[r.AWS_REGION]; ok {
				r.AMI_OWNER_ID = v
			}
		}
	}

	amiSearch(ctx, r)
}
