package cmd

import (
	"strconv"
	"strings"

	"github.com/guessi/eks-ami-finder/pkg/constants"
	"github.com/urfave/cli/v2"
)

func amiSearchInput(c *cli.Context) amiSearchInputSpec {
	maxResults, _ := strconv.Atoi(c.String("max-results"))

	return amiSearchInputSpec{
		AWS_REGION:         c.String("region"),
		AMI_OWNER_ID:       c.String("owner-id"),
		AMI_TYPE:           c.String("ami-type"),
		KUBERNETES_VERSION: c.String("kubernetes-version"),
		RELEASE_DATE:       c.String("release-date"),
		MAX_RESULTS:        maxResults,
		INCLUDE_DEPRECATED: c.Bool("include-deprecated"),
		DEBUG_MODE:         c.Bool("debug"),
	}
}

func Wrapper(c *cli.Context) {
	r := amiSearchInput(c)

	// if region input but having no ownerId assigned, assume it is looking for EKS official image build
	if len(r.AWS_REGION) > 0 && len(r.AMI_OWNER_ID) != 12 {
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

	amiSearch(r)
}
