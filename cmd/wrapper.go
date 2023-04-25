package cmd

import (
	"fmt"
	"time"

	"github.com/guessi/eks-ami-finder/pkg/constants"
	"github.com/urfave/cli/v2"
)

func Wrapper(c *cli.Context) {
	amiRegion := c.String("region")
	amiOwnerId := c.String("owner-id")
	amiType := c.String("ami-type")
	kubernetesVersion := c.String("kubernetes-version")
	releaseDate := c.String("release-date")
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

	// if release date string input is invalid, set current month as filter
	if len(releaseDate) <= 0 || len(releaseDate) > 8 {
		now := time.Now()
		releaseDate = fmt.Sprintf("%02d%02d", now.Year(), now.Month())
	}

	finder(
		amiRegion,
		amiOwnerId,
		amiType,
		kubernetesVersion,
		releaseDate,
		includeDeprecated,
		debugMode,
	)
}
