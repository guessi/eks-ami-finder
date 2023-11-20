package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/guessi/eks-ami-finder/pkg/constants"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func finder(region, ownerId, amiType, kubernetesVersion, releaseDate string, includeDeprecated, debug bool) {
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(region),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := ec2.NewFromConfig(cfg)

	pattern := fmt.Sprintf("%s-%s-v%s*", constants.AmiPrefixMappings[amiType], kubernetesVersion, releaseDate)

	filters := []types.Filter{
		types.Filter{
			Name: aws.String("owner-id"),
			Values: []string{
				ownerId,
			},
		},
		types.Filter{
			Name: aws.String("name"),
			Values: []string{
				pattern,
			},
		},
	}

	// NICE-TO-HAVE: should have pagination support
	amis, err := svc.DescribeImages(context.TODO(), &ec2.DescribeImagesInput{
		Filters:           filters,
		NextToken:         nil,
		IncludeDeprecated: aws.Bool(includeDeprecated),
	})
	if err != nil {
		var re *awshttp.ResponseError
		if errors.As(err, &re) {
			log.Fatalf("requestID: %s, error: %v", re.ServiceRequestID(), re.Unwrap())
		}
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{
		"Region",
		"AMI ID",
		"Name",
		"Description",
		"Creation Date",
		"DeprecationTime",
	})

	// tricky trick to sort AMI by creation date
	t.SortBy([]table.SortBy{{Name: "Creation Date", Mode: table.Dsc}})
	t.SetColumnConfigs([]table.ColumnConfig{{Name: "Creation Date", Hidden: true}})

	for _, image := range amis.Images {
		t.AppendRow(table.Row{region, *image.ImageId, *image.Name, *image.Description, *image.CreationDate, *image.DeprecationTime})
	}

	t.Style().Format.Header = text.FormatDefault
	t.Render()

	if debug {
		println()
		print(fmt.Sprintf("OwerId: %s\n", ownerId))
		print(fmt.Sprintf("Filter: %s\n", pattern))
	}

}
