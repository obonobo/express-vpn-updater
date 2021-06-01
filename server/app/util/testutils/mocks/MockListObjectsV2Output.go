package mocks

import (
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	date = time.Date(1997, time.July, 21, 0, 0, 0, 0, &time.Location{})
)

func MockListObjectsV2Output() *s3.ListObjectsV2Output {
	return &s3.ListObjectsV2Output{
		Contents: []*s3.Object{
			{
				Size:         aws.Int64(666),
				Key:          aws.String("my-package"),
				ETag:         aws.String("my-etag"),
				StorageClass: aws.String("my-storage-class"),
				Owner:        &s3.Owner{DisplayName: aws.String("me")},
				LastModified: &date,
			},
		},
	}
}
