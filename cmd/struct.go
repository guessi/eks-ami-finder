package cmd

type amiSearchInputSpec struct {
	AWS_REGION         string
	AMI_OWNER_ID       string
	AMI_TYPE           string
	KUBERNETES_VERSION string
	RELEASE_DATE       string
	MAX_RESULTS        int
	AUTO_MODE          bool
	INCLUDE_DEPRECATED bool
	DEBUG_MODE         bool
}
