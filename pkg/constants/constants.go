package constants

const (
	NAME  string = "eks-ami-finder"
	USAGE string = "Helper tool to find Amazon EKS optimized AMI IDs"
)

var (
	GitVersion string
	GoVersion  string
	BuildTime  string
)

var (
	// Valid AMI_TYPE definitions
	// - https://docs.aws.amazon.com/eks/latest/APIReference/API_Nodegroup.html#AmazonEKS-Type-Nodegroup-amiType
	ValidAmiTypes = map[string][]string{
		// General AMI Types for Amazon EKS
		"DEFAULT": {
			"AL2_ARM_64",
			"AL2_x86_64",
			"AL2_x86_64_GPU",
			"AL2023_ARM_64_NVIDIA",
			"AL2023_ARM_64_STANDARD",
			"AL2023_x86_64_NEURON",
			"AL2023_x86_64_NVIDIA",
			"AL2023_x86_64_STANDARD",
			"BOTTLEROCKET_ARM_64",
			"BOTTLEROCKET_ARM_64_FIPS",
			"BOTTLEROCKET_ARM_64_NVIDIA",
			"BOTTLEROCKET_x86_64",
			"BOTTLEROCKET_x86_64_FIPS",
			"BOTTLEROCKET_x86_64_NVIDIA",
			"WINDOWS_CORE_2016_x86_64",
			"WINDOWS_CORE_2019_x86_64",
			"WINDOWS_CORE_2022_x86_64",
			"WINDOWS_FULL_2016_x86_64",
			"WINDOWS_FULL_2019_x86_64",
			"WINDOWS_FULL_2022_x86_64",
		},
		// Special crafted AMI Types for Auto Mode
		"AUTO_MODE": {
			"AUTO_MODE_NEURON_x86_64",
			"AUTO_MODE_NVIDIA_ARM_64",
			"AUTO_MODE_NVIDIA_x86_64",
			"AUTO_MODE_STANDARD_ARM_64",
			"AUTO_MODE_STANDARD_x86_64",
		},
	}

	// Ideally, official AMI should comes from fixed AWS Account IDs, so hard-coded here should be fine.
	// Combine the output of GetParameter and pass it to DescribeImages, we can get fixed Account Id Mappings.
	// - https://docs.aws.amazon.com/eks/latest/userguide/retrieve-ami-id.html
	// - https://docs.aws.amazon.com/eks/latest/userguide/retrieve-windows-ami-id.html
	// - https://docs.aws.amazon.com/eks/latest/userguide/retrieve-ami-id-bottlerocket.html
	//
	// Regions introduced before March 20, 2019 are enabled by default, for the rest of others, "Opt-in" is required.
	// - https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-regions.html
	// - https://docs.aws.amazon.com/accounts/latest/reference/manage-acct-regions.html#manage-acct-regions-considerations
	AwsAccountMappingsAL = map[string]string{
		"af-south-1":     "877085696533",
		"ap-east-1":      "800184023465",
		"ap-east-2":      "533267051163",
		"ap-south-2":     "900889452093",
		"ap-southeast-3": "296578399912",
		"ap-southeast-4": "491585149902",
		"ap-southeast-5": "151610086707",
		"ap-southeast-6": "333609536671",
		"ap-southeast-7": "121268973566",
		"ca-west-1":      "761377655185",
		"cn-north-1":     "918309763551",
		"cn-northwest-1": "961992271922",
		"eu-central-2":   "900612956339",
		"eu-south-1":     "590381155156",
		"eu-south-2":     "455263428931",
		"il-central-1":   "066635153087",
		"me-central-1":   "759879836304",
		"mx-central-1":   "730335286997",
		"me-south-1":     "558608220178",
		"us-gov-east-1":  "151742754352",
		"us-gov-west-1":  "013241004608",
		"us-iso-east-1":  "725322719131",
		"us-isob-east-1": "187977181151",
		"*":              "602401143452",
	}

	AwsAccountMappingsBottlerocket = map[string]string{
		"af-south-1":     "291523557710",
		"ap-east-1":      "040063162771",
		"ap-east-2":      "992382384230",
		"ap-northeast-1": "593245189075",
		"ap-northeast-2": "630172235254",
		"ap-northeast-3": "429266927828",
		"ap-south-1":     "449901457613",
		"ap-south-2":     "713232077714",
		"ap-southeast-1": "406264879685",
		"ap-southeast-2": "823100568288",
		"ap-southeast-3": "888108400464",
		"ap-southeast-4": "445448481902",
		"ap-southeast-5": "191778180997",
		"ap-southeast-6": "933873600184",
		"ap-southeast-7": "058264547253",
		"ca-central-1":   "229026816814",
		"ca-west-1":      "458152013874",
		"cn-north-1":     "183365920950",
		"cn-northwest-1": "183890449532",
		"eu-central-1":   "149721548608",
		"eu-central-2":   "799456934533",
		"eu-north-1":     "432623269467",
		"eu-south-1":     "754205708698",
		"eu-south-2":     "082494185113",
		"eu-west-1":      "503807174151",
		"eu-west-2":      "941016683700",
		"eu-west-3":      "296779064547",
		"il-central-1":   "346625278781",
		"me-central-1":   "789853572315",
		"me-south-1":     "340903185543",
		"mx-central-1":   "481050282565",
		"sa-east-1":      "044060155884",
		"us-east-1":      "092701018921",
		"us-east-2":      "419346874475",
		"us-west-1":      "724952271658",
		"us-west-2":      "651937483462",
		"us-gov-east-1":  "372293620468",
		"us-gov-west-1":  "372333677703",
	}

	AwsAccountMappingsWindows = map[string]string{
		"af-south-1":     "597400817333",
		"ap-east-1":      "907716943349",
		"ap-east-2":      "891376989287",
		"ap-northeast-1": "517802777641",
		"ap-northeast-2": "918716859121",
		"ap-northeast-3": "834741791908",
		"ap-south-1":     "750252652416",
		"ap-south-2":     "063495277261",
		"ap-southeast-1": "954049747103",
		"ap-southeast-2": "687402702948",
		"ap-southeast-3": "086269339428",
		"ap-southeast-4": "591620260053",
		"ap-southeast-5": "250259214768",
		"ap-southeast-6": "550366911139",
		"ap-southeast-7": "730335552224",
		"ca-central-1":   "151453898909",
		"ca-west-1":      "492533705213",
		"cn-north-1":     "436023608783",
		"cn-northwest-1": "438848857437",
		"eu-central-1":   "999352223265",
		"eu-central-2":   "108439994008",
		"eu-north-1":     "142676981321",
		"eu-south-1":     "340946277815",
		"eu-south-2":     "992500858046",
		"eu-west-1":      "402743460324",
		"eu-west-2":      "111789216327",
		"eu-west-3":      "225872793654",
		"il-central-1":   "836514005396",
		"me-central-1":   "522606368197",
		"me-south-1":     "517341481761",
		"mx-central-1":   "931106322324",
		"sa-east-1":      "980913465755",
		"us-east-1":      "957547624766",
		"us-east-2":      "205223424851",
		"us-west-1":      "247341962726",
		"us-west-2":      "137057727718",
		"us-gov-east-1":  "055183720277",
		"us-gov-west-1":  "055189784373",
	}

	// Auto Mode have predefined AWS Accounts list
	// - https://docs.aws.amazon.com/eks/latest/userguide/auto-controls.html
	AwsAccountMappingsAutoMode = map[string]string{
		"af-south-1":     "471112993317",
		"ap-east-1":      "590183728416",
		"ap-east-2":      "381492200852",
		"ap-northeast-1": "851725346105",
		"ap-northeast-2": "992382805010",
		"ap-northeast-3": "891377407544",
		"ap-south-1":     "975049899075",
		"ap-south-2":     "590183737426",
		"ap-southeast-1": "339712723301",
		"ap-southeast-2": "058264376476",
		"ap-southeast-3": "471112941769",
		"ap-southeast-4": "590183863144",
		"ap-southeast-5": "654654202513",
		"ap-southeast-6": "905418310314",
		"ap-southeast-7": "533267217478",
		"ca-central-1":   "992382439851",
		"ca-west-1":      "767397959864",
		"eu-central-1":   "891376953411",
		"eu-central-2":   "381492036002",
		"eu-north-1":     "339712696471",
		"eu-south-1":     "975049955519",
		"eu-south-2":     "471112620929",
		"eu-west-1":      "381492008532",
		"eu-west-2":      "590184142468",
		"eu-west-3":      "891376969258",
		"il-central-1":   "590183797093",
		"me-central-1":   "637423494195",
		"me-south-1":     "905418070398",
		"mx-central-1":   "211125506622",
		"sa-east-1":      "339712709251",
		"us-east-1":      "992382739861",
		"us-east-2":      "975050179949",
		"us-west-1":      "975050035094",
		"us-west-2":      "767397842682",
	}
)
