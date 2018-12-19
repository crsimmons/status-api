package testsupport

import "github.com/aws/aws-sdk-go/service/s3"

// FakeS3Client implements s3 for testing
type FakeS3Client struct {
	FakeHeadBucket func(input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error)
}

// HeadBucket delegates to FakeHeadBucket which is dynamically set by the tests
func (client *FakeS3Client) HeadBucket(input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
	return client.FakeHeadBucket(input)
}
