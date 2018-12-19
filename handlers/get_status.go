package handlers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/crsimmons/status-api/iaas"
	"github.com/go-martini/martini"
)

// GetStatus retrieves status from a pre-defined S3 bucket
func GetStatus(resp http.ResponseWriter, req *http.Request, params martini.Params, s3Client iaas.Client) {
	bucket := os.Getenv("BUCKET")
	_, err := s3Client.HeadBucket(&s3.HeadBucketInput{Bucket: aws.String(bucket)})
	if err != nil {
		resp.WriteHeader(404)
		fmt.Fprintf(resp, "Invalid bucket %s", bucket)
		return
	}
	resp.WriteHeader(200)
}
