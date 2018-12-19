package iaas

import "github.com/aws/aws-sdk-go/service/s3"

type Client interface {
	HeadBucket(input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error)
}
