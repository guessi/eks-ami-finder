package cmd

import (
	"context"
	"errors"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"

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
		fmt.Println("Can not pass --max-results with a value lower or equal to 0.")
		os.Exit(1)
	}

	if !slices.Contains(constants.ValidAmiTypes, input.AMI_TYPE) {
		fmt.Printf("Invalid AMI_TYPE input (Valid input: %s)\n", strings.Join(constants.ValidAmiTypes, ", "))
		os.Exit(1)
	}

	if len(input.RELEASE_DATE) != 0 {
		// releaseDate is expected to have at least Year included.
		if len(input.RELEASE_DATE) < 4 || len(input.RELEASE_DATE) > 8 {
			fmt.Println("Invalid --release-date passed.")
			os.Exit(1)
		}

		// Amazon EKS was first released back at Jun 05, 2018
		// - https://aws.amazon.com/blogs/aws/amazon-eks-now-generally-available/
		if r, err := strconv.Atoi(input.RELEASE_DATE); err == nil {
			if r < 2018 {
				fmt.Println("Invalid --release-date passed.")
				os.Exit(1)
			}
		} else {
			fmt.Println("Invalid --release-date passed.")
			os.Exit(1)
		}
	}

	if len(input.RELEASE_DATE) != 0 && strings.HasPrefix(input.AMI_TYPE, "BOTTLEROCKET_") {
		fmt.Println("Bottlerocket don't support filter by release date")
		os.Exit(1)
	}

	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion(input.AWS_REGION),
	)
	if err != nil {
		fmt.Printf("unable to load SDK config, %v", err)
		os.Exit(1)
	}

	svc := ec2.NewFromConfig(cfg)

	var pattern string

	switch ami_type := input.AMI_TYPE; ami_type {
	case "AL2_ARM_64":
		pattern = fmt.Sprintf("amazon-eks-arm64-node-%s-v%s*", input.KUBERNETES_VERSION, input.RELEASE_DATE)
	case "AL2_x86_64":
		pattern = fmt.Sprintf("amazon-eks-node-%s-v%s*", input.KUBERNETES_VERSION, input.RELEASE_DATE)
	case "AL2_x86_64_GPU":
		pattern = fmt.Sprintf("amazon-eks-gpu-node-%s-v%s*", input.KUBERNETES_VERSION, input.RELEASE_DATE)
	case "AL2023_ARM_64_STANDARD":
		pattern = fmt.Sprintf("amazon-eks-node-al2023-arm64-standard-%s-v%s*", input.KUBERNETES_VERSION, input.RELEASE_DATE)
	case "AL2023_x86_64_STANDARD":
		pattern = fmt.Sprintf("amazon-eks-node-al2023-x86_64-standard-%s-v%s*", input.KUBERNETES_VERSION, input.RELEASE_DATE)
	case "BOTTLEROCKET_ARM_64":
		pattern = fmt.Sprintf("bottlerocket-aws-k8s-%s-aarch64-v*", input.KUBERNETES_VERSION)
	case "BOTTLEROCKET_x86_64":
		pattern = fmt.Sprintf("bottlerocket-aws-k8s-%s-x86_64-v*", input.KUBERNETES_VERSION)
	case "BOTTLEROCKET_ARM_64_NVIDIA":
		pattern = fmt.Sprintf("bottlerocket-aws-k8s-%s-nvidia-aarch64-v*", input.KUBERNETES_VERSION)
	case "BOTTLEROCKET_x86_64_NVIDIA":
		pattern = fmt.Sprintf("bottlerocket-aws-k8s-%s-nvidia-x86_64-v*", input.KUBERNETES_VERSION)
	case "WINDOWS_CORE_2019_x86_64":
		pattern = fmt.Sprintf("Windows_Server-2019-English-Core-EKS_Optimized-%s-%s*", input.KUBERNETES_VERSION, input.RELEASE_DATE)
	case "WINDOWS_FULL_2019_x86_64":
		pattern = fmt.Sprintf("Windows_Server-2019-English-Full-EKS_Optimized-%s-%s*", input.KUBERNETES_VERSION, input.RELEASE_DATE)
	case "WINDOWS_CORE_2022_x86_64":
		pattern = fmt.Sprintf("Windows_Server-2022-English-Core-EKS_Optimized-%s-%s*", input.KUBERNETES_VERSION, input.RELEASE_DATE)
	case "WINDOWS_FULL_2022_x86_64":
		pattern = fmt.Sprintf("Windows_Server-2022-English-Full-EKS_Optimized-%s-%s*", input.KUBERNETES_VERSION, input.RELEASE_DATE)
	default:
		fmt.Println("Invalid AMI_TYPE input")
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
			fmt.Printf("requestID: %s, error: %v", re.ServiceRequestID(), re.Unwrap())
			os.Exit(1)
		}
	}

	if len(images) == 0 {
		fmt.Println("No matching AMI found.")
		os.Exit(0)
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
