package main

import (
	"flag"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/yury-sannikov/go-cloudfront-deploy/s3tools"
)

// 1 - Add site under hosted zone
// .. domain nanme
// .. description
// ==> JSON with

var bucketName string
var region string

func init() {
	const (
		defaultBucketName = ""
		usageBucket       = "Amazon S3 bucket name to deploy"
		defaultRegionName = "us-west-1"
		usageRegion       = "Amazon S3 region"
	)
	flag.StringVar(&bucketName, "bucket_name", defaultBucketName, usageBucket)
	flag.StringVar(&bucketName, "b", defaultBucketName, usageBucket+" (shorthand)")
	flag.StringVar(&region, "region", defaultRegionName, usageRegion)
	flag.StringVar(&region, "r", defaultRegionName, usageRegion+" (shorthand)")
}

func main() {
	flag.Parse()
	if bucketName == "" {
		flag.PrintDefaults()
		return
	}
	fmt.Printf("Region: %s, Bucket Name: %s\n", region, bucketName)

	config := &aws.Config{Region: &region}
	sess := session.New(config)
	svc := s3.New(sess)

	s3tools.CheckOrCreateBucket(svc, bucketName)
}