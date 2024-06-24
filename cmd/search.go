package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/guessi/eks-ami-finder/pkg/constants"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func findAmiMatches(ctx context.Context, svc ec2.DescribeImagesAPIClient, input *ec2.DescribeImagesInput, maxResults int) ([]types.Image, error) {
	var images []types.Image
	var returnSize int

	paginator := ec2.NewDescribeImagesPaginator(svc, input)
	for paginator.HasMorePages() {
		out, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		images = append(images, out.Images...)
		if len(images) > maxResults {
			break
		}
	}

	returnSize = maxResults
	if len(images) < maxResults {
		returnSize = len(images)
	}
	return images[:returnSize], nil
}

func amiSearch(input amiSearchInputSpec) {
	// do nothing if maxResults is invalid input
	if input.MAX_RESULTS <= 0 {
		log.Fatalf("Can not pass --max-results with a value lower or equal to 0.\n")
	}

	if input.AMI_FAMILY == "AL2023" {
		if input.AMI_TYPE != "x86_64" && input.AMI_TYPE != "arm64" {
			log.Fatalf("Invalid combinition %s, %s\n", input.AMI_FAMILY, input.AMI_TYPE)
		}
	}

	if len(input.RELEASE_DATE) != 0 {
		// releaseDate is expected to have at least Year included.
		if len(input.RELEASE_DATE) < 4 || len(input.RELEASE_DATE) > 8 {
			log.Fatalf("Invalid --release-date passed.\n")
		}

		// Amazon EKS was first released back at Jun 05, 2018
		// - https://aws.amazon.com/blogs/aws/amazon-eks-now-generally-available/
		if r, err := strconv.Atoi(input.RELEASE_DATE); err == nil {
			if r < 2018 {
				log.Fatalf("Invalid --release-date passed.\n")
			}
		} else {
			log.Fatalf("Invalid --release-date passed.\n")
		}
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(input.AWS_REGION),
	)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	svc := ec2.NewFromConfig(cfg)

	pattern := fmt.Sprintf("%s-%s-v%s*", constants.AmiPrefixMappings[input.AMI_TYPE], input.KUBERNETES_VERSION, input.RELEASE_DATE)
	if input.AMI_FAMILY == "AL2023" {
		pattern = fmt.Sprintf("%s-%s-v%s*", constants.AmiPrefixMappingsAL2023[input.AMI_TYPE], input.KUBERNETES_VERSION, input.RELEASE_DATE)
	}

	filters := []types.Filter{
		{
			Name: aws.String("owner-id"),
			Values: []string{
				input.AMI_OWNER_ID,
			},
		},
		{
			Name: aws.String("name"),
			Values: []string{
				pattern,
			},
		},
	}

	describeImagesInput := ec2.DescribeImagesInput{
		Filters:           filters,
		NextToken:         nil,
		IncludeDeprecated: aws.Bool(input.INCLUDE_DEPRECATED),
	}

	images, err := findAmiMatches(context.TODO(), svc, &describeImagesInput, input.MAX_RESULTS)
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
		"Architecture",
	})

	// tricky trick to sort AMI by creation date
	t.SortBy([]table.SortBy{{Name: "Creation Date", Mode: table.Dsc}})
	t.SetColumnConfigs([]table.ColumnConfig{{Name: "Creation Date", Hidden: true}})

	for _, i := range images {
		t.AppendRow(table.Row{
			input.AWS_REGION,
			*i.ImageId,
			*i.Name,
			*i.Description,
			*i.CreationDate,
			*i.DeprecationTime,
			i.Architecture,
		})
	}

	t.Style().Format.Header = text.FormatDefault
	t.Render()

	if input.DEBUG_MODE {
		println()
		print(fmt.Sprintf("OwerId: %s\n", input.AMI_OWNER_ID))
		print(fmt.Sprintf("Filter: %s\n", pattern))
	}

}
