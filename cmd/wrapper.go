package cmd

import (
	"strconv"

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

	// if region input but having no ownerId assigned, assume it is loolking for EKS official image build
	if len(r.AWS_REGION) > 0 && len(r.AMI_OWNER_ID) != 12 {
		r.AMI_OWNER_ID = constants.AwsAccountMappings["*"]

		for k, v := range constants.AwsAccountMappings {
			if k == r.AWS_REGION {
				r.AMI_OWNER_ID = v
			}
		}
	}

	amiSearch(r)
}
