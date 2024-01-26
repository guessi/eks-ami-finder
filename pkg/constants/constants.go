package constants

const (
	NAME    string = "eks-ami-finder"
	USAGE   string = "retrieve Amazon EKS AMI with filters"
	VERSION string = "1.0.7"
)

var (
	AmiPrefixMappings = map[string]string{
		"arm64":      "amazon-eks-arm64-node",
		"x86_64":     "amazon-eks-node",
		"x86_64-gpu": "amazon-eks-gpu-node",
	}

	AwsAccountMappings = map[string]string{
		"af-south-1":     "877085696533",
		"ap-east-1":      "800184023465",
		"ap-south-2":     "900889452093",
		"ap-southeast-3": "296578399912",
		"ap-southeast-4": "491585149902",
		"ca-west-1":      "761377655185",
		"cn-north-1":     "918309763551",
		"cn-northwest-1": "961992271922",
		"eu-central-2":   "900612956339",
		"eu-south-1":     "590381155156",
		"eu-south-2":     "455263428931",
		"il-central-1":   "066635153087",
		"me-central-1":   "759879836304",
		"me-south-1":     "558608220178",
		"us-gov-east-1":  "151742754352",
		"us-gov-west-1":  "013241004608",
		"us-iso-east-1":  "725322719131",
		"us-isob-east-1": "187977181151",
		"*":              "602401143452",
	}
)
