package cmd

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	awshttp "github.com/aws/aws-sdk-go-v2/aws/transport/http"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"github.com/aws/aws-sdk-go-v2/service/ec2/types"
	"github.com/guessi/eks-ami-finder/pkg/constants"
	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

func isUnsupportedRegion(ctx context.Context, region string) bool {
	if region == "" {
		return true
	}

	hostname := fmt.Sprintf("ec2.%s.amazonaws.com", region)
	lookupCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	_, err := (&net.Resolver{}).LookupHost(lookupCtx, hostname)
	return err != nil
}

func findAmiMatches(ctx context.Context, svc ec2.DescribeImagesAPIClient, input *ec2.DescribeImagesInput, maxResults int) ([]types.Image, error) {
	var images []types.Image
	var returnSize int

	paginator := ec2.NewDescribeImagesPaginator(svc, input)
	for paginator.HasMorePages() {
		// Check for context cancellation
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		out, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}
		images = append(images, out.Images...)
		if len(images) > maxResults {
			break
		}
	}

	returnSize = min(maxResults, len(images))

	return images[:returnSize], nil
}

func amiSearch(ctx context.Context, input amiSearchInputSpec) {
	// do nothing if maxResults is invalid input
	if input.MAX_RESULTS <= 0 {
		fmt.Printf("Can not pass --max-results with a value lower or equal to 0.\n\n")
		os.Exit(1)
	}

	if isUnsupportedRegion(ctx, input.AWS_REGION) {
		fmt.Printf("Unable to resolve EC2 endpoint for the given region. Please check your region input.\n\n")
		os.Exit(1)
	}

	var majorK8sVersion, minorK8sVersion int
	versionParts := strings.Split(input.KUBERNETES_VERSION, ".")
	if len(versionParts) >= 2 {
		if i, err := strconv.Atoi(versionParts[0]); err == nil {
			majorK8sVersion = int(i)
		}
		if i, err := strconv.Atoi(versionParts[1]); err == nil {
			minorK8sVersion = int(i)
		}
	} else {
		fmt.Printf("Invalid Kubernetes version format. Expected format: X.Y (e.g., 1.33)\n\n")
		os.Exit(1)
	}

	if majorK8sVersion != 1 || minorK8sVersion < 10 {
		// the first Amazon EKS version was 1.10
		// - https://aws.amazon.com/blogs/aws/amazon-eks-now-generally-available/
		fmt.Printf("The very first Amazon EKS version was 1.10, so there would have no Amazon EKS %s.\n\n", input.KUBERNETES_VERSION)
		os.Exit(1)
	}

	if !slices.Contains(constants.ValidAmiTypes, input.AMI_TYPE) {
		fmt.Printf("Invalid AMI_TYPE input (Valid input: %s).\n\n", strings.Join(constants.ValidAmiTypes, ", "))
		os.Exit(1)
	} else {
		// AL2 AMI will no longer supported for Amazon EKS 1.33 or newer
		// - https://docs.aws.amazon.com/eks/latest/userguide/eks-ami-deprecation-faqs.html
		if minorK8sVersion >= 33 && strings.Split(input.AMI_TYPE, "_")[0] == "AL2" {
			fmt.Printf("There would have no AL2-based Optimized AMI support for Amazon EKS 1.33 or newer.\n\n")
			fmt.Printf("Check the following link for more info:\n")
			fmt.Printf("- https://docs.aws.amazon.com/eks/latest/userguide/eks-ami-deprecation-faqs.html\n\n")
			os.Exit(1)
		}

		// AL2023 AMI support starting from Amazon EKS 1.23 or newer
		// - https://aws.amazon.com/blogs/containers/amazon-eks-optimized-amazon-linux-2023-amis-now-available/
		// - https://aws.amazon.com/blogs/containers/amazon-eks-optimized-amazon-linux-2023-accelerated-amis-now-available/
		// - https://docs.aws.amazon.com/eks/latest/userguide/doc-history.html
		if minorK8sVersion < 23 && strings.Split(input.AMI_TYPE, "_")[0] == "AL2023" {
			fmt.Printf("Invalid input: there have no %s support for Amazon EKS %s.\n\n", input.AMI_TYPE, input.KUBERNETES_VERSION)
			os.Exit(1)
		}

		// Bottlerocket AMI initially support Amazon EKS 1.15 or newer
		// - https://aws.amazon.com/blogs/containers/amazon-eks-adds-native-support-for-bottlerocket-in-managed-node-groups/
		// - https://github.com/bottlerocket-os/bottlerocket/releases/tag/v1.0.0
		// - https://docs.aws.amazon.com/eks/latest/userguide/doc-history.html
		if minorK8sVersion < 15 && strings.Split(input.AMI_TYPE, "_")[0] == "BOTTLEROCKET" {
			fmt.Printf("Invalid input: there have no %s support for Amazon EKS %s.\n\n", input.AMI_TYPE, input.KUBERNETES_VERSION)
			os.Exit(1)
		}

		// Windows Server AMI initially support Amazon EKS 1.14 or newer
		// - https://docs.aws.amazon.com/eks/latest/userguide/doc-history.html
		// - https://github.com/aws/containers-roadmap/issues/69#issuecomment-539641916
		// - https://docs.aws.amazon.com/eks/latest/userguide/doc-history.html
		if minorK8sVersion < 14 && strings.Split(input.AMI_TYPE, "_")[0] == "WINDOWS" {
			fmt.Printf("Invalid input: there have no %s support for Amazon EKS %s.\n\n", input.AMI_TYPE, input.KUBERNETES_VERSION)
			os.Exit(1)
		}

		// Windows Server 2019/2022 only support Amazon EKS 1.23 or newer
		// - https://aws.amazon.com/blogs/containers/deploying-amazon-eks-windows-managed-node-groups/
		// - https://docs.aws.amazon.com/eks/latest/userguide/doc-history.html
		if minorK8sVersion < 23 && strings.Split(input.AMI_TYPE, "_")[0] == "WINDOWS" {
			amiTypeParts := strings.Split(input.AMI_TYPE, "_")
			if len(amiTypeParts) >= 3 && (amiTypeParts[2] == "2019" || amiTypeParts[2] == "2022") {
				fmt.Printf("Invalid input: there have no %s support for Amazon EKS %s.\n\n", input.AMI_TYPE, input.KUBERNETES_VERSION)
				os.Exit(1)
			}
		}

	}

	if releaseDate := input.RELEASE_DATE; len(releaseDate) != 0 {
		// releaseDate is expected to have at least Year included.
		if len(releaseDate) < 4 || len(releaseDate) > 8 {
			fmt.Printf("Invalid --release-date passed.\n\n")
			os.Exit(1)
		}

		// Amazon EKS was first released back at Jun 05, 2018
		// - https://aws.amazon.com/blogs/aws/amazon-eks-now-generally-available/
		if year, err := strconv.Atoi(releaseDate[:4]); err != nil || year < 2018 {
			fmt.Printf("Invalid --release-date passed.\n\n")
			os.Exit(1)
		}

		// Bottlerocket AMIs don't support release date filtering
		if strings.HasPrefix(input.AMI_TYPE, "BOTTLEROCKET_") {
			fmt.Printf("Bottlerocket doesn't support filter by release date.\n\n")
			os.Exit(1)
		}
	}

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(input.AWS_REGION),
	)
	if err != nil {
		fmt.Printf("unable to load SDK config, %v\n\n", err)
		os.Exit(1)
	}

	svc := ec2.NewFromConfig(cfg)

	var pattern string

	// Map of AMI type patterns to avoid repetitive string formatting
	amiPatterns := map[string]string{
		"AL2_ARM_64":                 "amazon-eks-arm64-node-%s-v%s*",
		"AL2_x86_64_GPU":             "amazon-eks-gpu-node-%s-v%s*",
		"AL2_x86_64":                 "amazon-eks-node-%s-v%s*",
		"AL2023_ARM_64_NVIDIA":       "amazon-eks-node-al2023-arm64-nvidia-%s-v%s*",
		"AL2023_ARM_64_STANDARD":     "amazon-eks-node-al2023-arm64-standard-%s-v%s*",
		"AL2023_x86_64_NEURON":       "amazon-eks-node-al2023-x86_64-neuron-%s-v%s*",
		"AL2023_x86_64_NVIDIA":       "amazon-eks-node-al2023-x86_64-nvidia-%s-v%s*",
		"AL2023_x86_64_STANDARD":     "amazon-eks-node-al2023-x86_64-standard-%s-v%s*",
		"BOTTLEROCKET_ARM_64_FIPS":   "bottlerocket-aws-k8s-%s-fips-aarch64-v*",
		"BOTTLEROCKET_ARM_64_NVIDIA": "bottlerocket-aws-k8s-%s-nvidia-aarch64-v*",
		"BOTTLEROCKET_ARM_64":        "bottlerocket-aws-k8s-%s-aarch64-v*",
		"BOTTLEROCKET_x86_64_FIPS":   "bottlerocket-aws-k8s-%s-fips-x86_64-v*",
		"BOTTLEROCKET_x86_64_NVIDIA": "bottlerocket-aws-k8s-%s-nvidia-x86_64-v*",
		"BOTTLEROCKET_x86_64":        "bottlerocket-aws-k8s-%s-x86_64-v*",
		"WINDOWS_CORE_2016_x86_64":   "Windows_Server-2016-English-Core-EKS_Optimized-%s-%s*",
		"WINDOWS_CORE_2019_x86_64":   "Windows_Server-2019-English-Core-EKS_Optimized-%s-%s*",
		"WINDOWS_CORE_2022_x86_64":   "Windows_Server-2022-English-Core-EKS_Optimized-%s-%s*",
		"WINDOWS_FULL_2016_x86_64":   "Windows_Server-2016-English-Full-EKS_Optimized-%s-%s*",
		"WINDOWS_FULL_2019_x86_64":   "Windows_Server-2019-English-Full-EKS_Optimized-%s-%s*",
		"WINDOWS_FULL_2022_x86_64":   "Windows_Server-2022-English-Full-EKS_Optimized-%s-%s*",
	}

	if patternTemplate, ok := amiPatterns[input.AMI_TYPE]; ok {
		if strings.HasPrefix(input.AMI_TYPE, "BOTTLEROCKET_") {
			pattern = fmt.Sprintf(patternTemplate, input.KUBERNETES_VERSION)
		} else {
			pattern = fmt.Sprintf(patternTemplate, input.KUBERNETES_VERSION, input.RELEASE_DATE)
		}
	} else {
		fmt.Printf("Invalid AMI_TYPE input.\n\n")
		os.Exit(1)
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

	images, err := findAmiMatches(ctx, svc, &describeImagesInput, input.MAX_RESULTS)
	if err != nil {
		// Check for context cancellation first
		if ctx.Err() != nil {
			fmt.Printf("Request cancelled or timed out: %v\n\n", ctx.Err())
			os.Exit(1)
		}

		// Check for AWS-specific errors
		var re *awshttp.ResponseError
		if errors.As(err, &re) {
			fmt.Printf("requestID: %s, error: %v\n\n", re.ServiceRequestID(), re.Unwrap())
			os.Exit(1)
		}

		// Handle other errors
		fmt.Printf("Error retrieving AMI information: %v\n\n", err)
		os.Exit(1)
	}

	if len(images) == 0 {
		fmt.Printf("No matching AMI found.\n\n")
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
			aws.ToString(i.ImageId),
			aws.ToString(i.Name),
			aws.ToString(i.Description),
			aws.ToString(i.CreationDate),
			aws.ToString(i.DeprecationTime),
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
