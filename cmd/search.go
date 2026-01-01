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

	// Always fetch at least 50 AMIs to ensure we get recent ones after sorting
	fetchLimit := max(maxResults*2, 50)

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
		if len(images) > fetchLimit {
			break
		}
	}

	returnSize = min(maxResults, len(images))

	return images[:returnSize], nil
}

func simpleInputValidation(ctx context.Context, input amiSearchInputSpec) error {
	if isUnsupportedRegion(ctx, input.AWS_REGION) {
		return fmt.Errorf("unable to resolve EC2 endpoint for the given region. Please check your region input")
	}

	// Parse Kubernetes version for cross-validation
	versionParts := strings.Split(input.KUBERNETES_VERSION, ".")
	var minorK8sVersion int
	if len(versionParts) == 2 {
		if i, err := strconv.Atoi(versionParts[1]); err == nil {
			minorK8sVersion = i
		}
	}

	// Auto Mode validation
	if input.AUTO_MODE {
		if !slices.Contains(constants.ValidAmiTypes["AUTO_MODE"], input.AMI_TYPE) {
			return fmt.Errorf("invalid --ami-type input for auto-mode (Valid input: %s)", strings.Join(constants.ValidAmiTypes["AUTO_MODE"], ", "))
		}

		// Auto Mode only available for Amazon EKS 1.29 or later
		// - https://docs.aws.amazon.com/eks/latest/userguide/create-auto.html
		if minorK8sVersion < 29 {
			return fmt.Errorf("EKS Auto Mode requires Kubernetes version 1.29 or greater. See: https://docs.aws.amazon.com/eks/latest/userguide/create-auto.html")
		}
	} else {
		if !slices.Contains(constants.ValidAmiTypes["DEFAULT"], input.AMI_TYPE) {
			return fmt.Errorf("invalid --ami-type input (Valid input: %s)", strings.Join(constants.ValidAmiTypes["DEFAULT"], ", "))
		}

		// AMI type specific validations
		if strings.HasPrefix(input.AMI_TYPE, "AL2_") {
			// AL2 AMI will no longer be supported for Amazon EKS 1.33 or newer
			// - https://docs.aws.amazon.com/eks/latest/userguide/eks-ami-deprecation-faqs.html
			if minorK8sVersion >= 33 {
				return fmt.Errorf("AL2-based AMI is not supported for Amazon EKS 1.33 or newer. See: https://docs.aws.amazon.com/eks/latest/userguide/eks-ami-deprecation-faqs.html")
			}
		}

		if strings.HasPrefix(input.AMI_TYPE, "AL2023_") {
			// AL2023 AMI support starting from Amazon EKS 1.23 or newer
			// - https://aws.amazon.com/blogs/containers/amazon-eks-optimized-amazon-linux-2023-amis-now-available/
			// - https://aws.amazon.com/blogs/containers/amazon-eks-optimized-amazon-linux-2023-accelerated-amis-now-available/
			// - https://docs.aws.amazon.com/eks/latest/userguide/doc-history.html
			if minorK8sVersion < 23 {
				return fmt.Errorf("%s requires Amazon EKS 1.23 or newer (you specified %s)", input.AMI_TYPE, input.KUBERNETES_VERSION)
			}
		}

		if strings.HasPrefix(input.AMI_TYPE, "BOTTLEROCKET_") {
			// BOTTLEROCKET NVIDIA FIPS variants only available for Kubernetes 1.29+
			// - https://github.com/bottlerocket-os/bottlerocket/releases/tag/v1.51.0
			// - https://github.com/bottlerocket-os/bottlerocket/pull/4671
			if minorK8sVersion < 29 && strings.HasSuffix(input.AMI_TYPE, "NVIDIA_FIPS") {
				return fmt.Errorf("%s requires Amazon EKS 1.29 or newer (you specified %s)", input.AMI_TYPE, input.KUBERNETES_VERSION)
			}
			// Bottlerocket AMI initially support Amazon EKS 1.15 or newer
			// - https://aws.amazon.com/blogs/containers/amazon-eks-adds-native-support-for-bottlerocket-in-managed-node-groups/
			// - https://github.com/bottlerocket-os/bottlerocket/releases/tag/v1.0.0
			// - https://docs.aws.amazon.com/eks/latest/userguide/doc-history.html
			if minorK8sVersion < 15 {
				return fmt.Errorf("%s requires Amazon EKS 1.15 or newer (you specified %s)", input.AMI_TYPE, input.KUBERNETES_VERSION)
			}
		}

		if strings.HasPrefix(input.AMI_TYPE, "WINDOWS_") {
			// Windows Server 2019/2022 only support Amazon EKS 1.23 or newer
			// - https://aws.amazon.com/blogs/containers/deploying-amazon-eks-windows-managed-node-groups/
			// - https://docs.aws.amazon.com/eks/latest/userguide/doc-history.html
			if minorK8sVersion < 23 {
				amiTypeParts := strings.Split(input.AMI_TYPE, "_")
				if len(amiTypeParts) >= 3 && (amiTypeParts[2] == "2019" || amiTypeParts[2] == "2022") {
					return fmt.Errorf("%s requires Amazon EKS 1.23 or newer (you specified %s)", input.AMI_TYPE, input.KUBERNETES_VERSION)
				}
			}
			// Windows Server AMI initially support Amazon EKS 1.14 or newer
			// - https://docs.aws.amazon.com/eks/latest/userguide/doc-history.html
			// - https://github.com/aws/containers-roadmap/issues/69#issuecomment-539641916
			if minorK8sVersion < 14 {
				return fmt.Errorf("%s requires Amazon EKS 1.14 or newer (you specified %s)", input.AMI_TYPE, input.KUBERNETES_VERSION)
			}
		}
	}

	return nil
}

func amiSearch(ctx context.Context, input amiSearchInputSpec) error {
	// basic validations
	if err := simpleInputValidation(ctx, input); err != nil {
		return err
	}

	// Additional release date validation (requires AMI type context)
	if releaseDate := input.RELEASE_DATE; len(releaseDate) != 0 {
		// Amazon EKS was first released back at Jun 05, 2018
		// - https://aws.amazon.com/blogs/aws/amazon-eks-now-generally-available/
		if year, err := strconv.Atoi(releaseDate[:4]); err != nil || year < 2018 {
			return fmt.Errorf("invalid release-date. Amazon EKS was first released in 2018")
		}

		// Bottlerocket AMIs don't support release date filtering
		if !input.AUTO_MODE && strings.HasPrefix(input.AMI_TYPE, "BOTTLEROCKET_") {
			return fmt.Errorf("Bottlerocket doesn't support filter by release date")
		}
	}

	cfg, err := config.LoadDefaultConfig(
		ctx,
		config.WithRegion(input.AWS_REGION),
	)
	if err != nil {
		return fmt.Errorf("unable to load SDK config: %v", err)
	}

	svc := ec2.NewFromConfig(cfg)

	var pattern string

	// Map of AMI type patterns to avoid repetitive string formatting (Auto Mode)
	autoModeAmiPatterns := map[string]string{
		"AUTO_MODE_NEURON_x86_64":   "eks-auto-neuron-%s-x86_64-%s*",
		"AUTO_MODE_NVIDIA_ARM_64":   "eks-auto-nvidia-%s-aarch64-%s*",
		"AUTO_MODE_NVIDIA_x86_64":   "eks-auto-nvidia-%s-x86_64-%s*",
		"AUTO_MODE_STANDARD_ARM_64": "eks-auto-standard-%s-aarch64-%s*",
		"AUTO_MODE_STANDARD_x86_64": "eks-auto-standard-%s-x86_64-%s*",
	}

	// Map of AMI type patterns to avoid repetitive string formatting
	amiPatterns := map[string]string{
		"AL2_ARM_64":                      "amazon-eks-arm64-node-%s-v%s*",
		"AL2_x86_64_GPU":                  "amazon-eks-gpu-node-%s-v%s*",
		"AL2_x86_64":                      "amazon-eks-node-%s-v%s*",
		"AL2023_ARM_64_NVIDIA":            "amazon-eks-node-al2023-arm64-nvidia-%s-v%s*",
		"AL2023_ARM_64_STANDARD":          "amazon-eks-node-al2023-arm64-standard-%s-v%s*",
		"AL2023_x86_64_NEURON":            "amazon-eks-node-al2023-x86_64-neuron-%s-v%s*",
		"AL2023_x86_64_NVIDIA":            "amazon-eks-node-al2023-x86_64-nvidia-%s-v%s*",
		"AL2023_x86_64_STANDARD":          "amazon-eks-node-al2023-x86_64-standard-%s-v%s*",
		"BOTTLEROCKET_ARM_64_FIPS":        "bottlerocket-aws-k8s-%s-fips-aarch64-v*",
		"BOTTLEROCKET_ARM_64_NVIDIA":      "bottlerocket-aws-k8s-%s-nvidia-aarch64-v*",
		"BOTTLEROCKET_ARM_64_NVIDIA_FIPS": "bottlerocket-aws-k8s-%s-nvidia-fips-aarch64-v*",
		"BOTTLEROCKET_ARM_64":             "bottlerocket-aws-k8s-%s-aarch64-v*",
		"BOTTLEROCKET_x86_64_FIPS":        "bottlerocket-aws-k8s-%s-fips-x86_64-v*",
		"BOTTLEROCKET_x86_64_NVIDIA":      "bottlerocket-aws-k8s-%s-nvidia-x86_64-v*",
		"BOTTLEROCKET_x86_64_NVIDIA_FIPS": "bottlerocket-aws-k8s-%s-nvidia-fips-x86_64-v*",
		"BOTTLEROCKET_x86_64":             "bottlerocket-aws-k8s-%s-x86_64-v*",
		"WINDOWS_CORE_2016_x86_64":        "Windows_Server-2016-English-Core-EKS_Optimized-%s-%s*",
		"WINDOWS_CORE_2019_x86_64":        "Windows_Server-2019-English-Core-EKS_Optimized-%s-%s*",
		"WINDOWS_CORE_2022_x86_64":        "Windows_Server-2022-English-Core-EKS_Optimized-%s-%s*",
		"WINDOWS_FULL_2016_x86_64":        "Windows_Server-2016-English-Full-EKS_Optimized-%s-%s*",
		"WINDOWS_FULL_2019_x86_64":        "Windows_Server-2019-English-Full-EKS_Optimized-%s-%s*",
		"WINDOWS_FULL_2022_x86_64":        "Windows_Server-2022-English-Full-EKS_Optimized-%s-%s*",
	}

	if input.AUTO_MODE {
		if v, ok := constants.AwsAccountMappingsAutoMode[input.AWS_REGION]; ok {
			input.AMI_OWNER_ID = v
		} else {
			return fmt.Errorf("Auto Mode might not be supported in %s region", input.AWS_REGION)
		}

		if patternTemplate, ok := autoModeAmiPatterns[input.AMI_TYPE]; ok {
			pattern = fmt.Sprintf(patternTemplate, input.KUBERNETES_VERSION, input.RELEASE_DATE)
		} else {
			return fmt.Errorf("invalid ami-type input: %s", input.AMI_TYPE)
		}
	} else {
		if patternTemplate, ok := amiPatterns[input.AMI_TYPE]; ok {
			if strings.HasPrefix(input.AMI_TYPE, "BOTTLEROCKET_") {
				pattern = fmt.Sprintf(patternTemplate, input.KUBERNETES_VERSION)
			} else {
				pattern = fmt.Sprintf(patternTemplate, input.KUBERNETES_VERSION, input.RELEASE_DATE)
			}
		} else {
			return fmt.Errorf("invalid ami-type input: %s", input.AMI_TYPE)
		}
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
			return fmt.Errorf("request cancelled or timed out: %v", ctx.Err())
		}

		// Check for AWS-specific errors
		var re *awshttp.ResponseError
		if errors.As(err, &re) {
			return fmt.Errorf("AWS error (requestID: %s): %v", re.ServiceRequestID(), re.Unwrap())
		}

		return fmt.Errorf("error retrieving AMI information: %v", err)
	}

	if len(images) == 0 {
		fmt.Printf("No matching AMI found.\n\n")
		return nil
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

	return nil
}
