package cmd

import (
	"strconv"

	"github.com/guessi/eks-ami-finder/pkg/constants"
	"github.com/urfave/cli/v2"
)

func Wrapper(c *cli.Context) {
	amiRegion := c.String("region")
	amiOwnerId := c.String("owner-id")
	amiType := c.String("ami-type")
	kubernetesVersion := c.String("kubernetes-version")
	releaseDate := c.String("release-date")
	maxResults, _ := strconv.Atoi(c.String("max-results"))
	includeDeprecated := c.Bool("include-deprecated")
	debugMode := c.Bool("debug")

	// if region input but having no ownerId assigned, assume it is loolking for EKS official image build
	if len(amiRegion) > 0 && len(amiOwnerId) != 12 {
		amiOwnerId = constants.AwsAccountMappings["*"]

		for k, v := range constants.AwsAccountMappings {
			if k == amiRegion {
				amiOwnerId = v
			}
		}
	}

	finder(
		amiRegion,
		amiOwnerId,
		amiType,
		kubernetesVersion,
		releaseDate,
		maxResults,
		includeDeprecated,
		debugMode,
	)
}
