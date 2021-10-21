package common

const (
	PipeFileName = "/vxworks/comms/input.out"
	BucketOutputName = "vandv-common-store"
	BucketOutputFolder = "test-job-logs"
	OutputFilterRegex = `(?s)((Test|vxTest)\sOptions\sSummary).*?(Test\sexecution\sfinished)`
	EofFilter = "revoir"
	ClosePipeFileSeparator = ""
	JobParamDefaultValue = "ebf0001"
	AwsRegionEnvKey = "AWS_REGION"

	FilePermissions = 0666
	PidParamDefaultValue = 0
	SysKillSignal = 15
)
