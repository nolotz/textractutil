package textractutil

import "github.com/aws/aws-sdk-go-v2/service/textract/types"

type (
	AsyncOutput struct {
		JobId            string
		Status           types.JobStatus
		API              string
		JobTag           string
		Timestamp        int
		DocumentLocation AsyncDocumentLocation
	}

	AsyncDocumentLocation struct {
		S3ObjectName string
		S3Bucket     string
	}
)
