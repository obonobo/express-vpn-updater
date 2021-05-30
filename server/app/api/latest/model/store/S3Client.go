package store

import "github.com/aws/aws-sdk-go/service/s3"

type S3Client interface {

	// List the objects in a bucket
	ListObjectsV2(*s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)

	// Place an object into S3
	PutObject(*s3.PutObjectInput) (*s3.PutObjectOutput, error)
}
