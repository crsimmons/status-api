package handlers_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-martini/martini"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/crsimmons/status-api/handlers"
	"github.com/crsimmons/status-api/testsupport"
)

var (
	rr       *httptest.ResponseRecorder
	req      *http.Request
	s3Client *testsupport.FakeS3Client
	params   martini.Params
)

var _ = Describe("Handlers", func() {
	Describe("#GetHealth", func() {
		BeforeEach(func() {
			rr = httptest.NewRecorder()
			req = httptest.NewRequest("GET", "/v0/health", nil)
			GetHealth(rr, req)
		})
		Context("When the app is healthy", func() {
			It("Returns a 200 status code", func() {
				Expect(rr.Code).To(Equal(200))
			})
		})
	})

	Describe("#GetStatus", func() {
		BeforeEach(func() {
			rr = httptest.NewRecorder()
			s3Client = &testsupport.FakeS3Client{
				FakeHeadBucket: func(input *s3.HeadBucketInput) (*s3.HeadBucketOutput, error) {
					if *input.Bucket == "valid-bucket" {
						return nil, nil
					} else {
						return nil, fmt.Errorf("invalid bucket")
					}
				},
			}
		})
		Context("When the bucket does not exist", func() {
			JustBeforeEach(func() {
				os.Setenv("BUCKET", "invalid-bucket")
				req = httptest.NewRequest("GET", "/v0/valid-env/status", nil)
				params = make(map[string]string)
				params["env"] = "environment"
				GetStatus(rr, req, params, s3Client)
			})
			It("returns an appropriate error", func() {
				Expect(rr.Code).To(Equal(404))
				Expect(rr.Body.String()).To(Equal("Invalid bucket invalid-bucket"))
			})
		})
	})
})
