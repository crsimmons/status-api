package main

import (
	"fmt"
	"net/http"
	"os"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/crsimmons/status-api/handlers"
	"github.com/go-martini/martini"
)

func main() {
	version := os.Getenv("API_VERSION")
	if version == "" {
		version = "v0"
	}
	region := os.Getenv("REGION")
	endpoint := os.Getenv("ENDPOINT")


	healthEndpoint := fmt.Sprintf("/%s/health", version)
	statusEndpoint := fmt.Sprintf("/%s/:env/status", version)

	m := martini.Classic()
	m.Get(healthEndpoint, func(resp http.ResponseWriter, req *http.Request) {
		handlers.GetHealth(resp, req)
	})
	m.Get(statusEndpoint, func(resp http.ResponseWriter, req *http.Request, params martini.Params) {
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(region),
			Endpoint: aws.String(endpoint),
		})
		if err != nil {
			log.Fatal(err)
		}
		s3Client := s3.New(sess)
		handlers.GetStatus(resp, req, params, s3Client)
	})
	m.Run()
}
