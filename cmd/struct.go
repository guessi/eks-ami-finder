package cmd

type amiSearchInputSpec struct {
	AWS_REGION         string
	AMI_OWNER_ID       string
	AMI_TYPE           string
	AMI_FAMILY         string
	KUBERNETES_VERSION string
	RELEASE_DATE       string
	MAX_RESULTS        int
	INCLUDE_DEPRECATED bool
	DEBUG_MODE         bool
}
