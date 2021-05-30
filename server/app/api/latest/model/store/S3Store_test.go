package store

import (
	"net/http"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/obonobo/express-vpn-updater/server/app/util/testutils/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	testBucket = "test-bucket"
)

var (
	date = time.Date(1997, time.July, 21, 0, 0, 0, 0, &time.Location{})

	mockListObjectsV2Output = &s3.ListObjectsV2Output{
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

	mockPutObjectOutput = &s3.PutObjectOutput{}
)

// Tests S3Store.Get()
func TestGet(t *testing.T) {
	mockHttpClient := createMockHttpClient()
	mockS3Client := createMockS3Client()

	store := New(mockS3Client, mockHttpClient, testBucket)
	resp, err := store.Get()

	logger.Println(resp, err)

	if err != nil {
		assert.Fail(t, "g")
	}

	assert.True(t, true, "TODO")
}

// Tests S3Store.Put()
func TestPut(t *testing.T) {
	assert.True(t, true, "TODO")
}

func createMockS3Client() *mocks.MockS3Client {
	return mocks.
		NewMockS3Client().
		WithListObjectsV2(func(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
			return mockListObjectsV2Output, nil
		}).
		WithPutObject(func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			return mockPutObjectOutput, nil
		})
}

func createMockHttpClient() *mocks.MockHttpClient {
	return mocks.
		NewMockHttpClient().
		WithGet(func(url string) (*http.Response, error) {
			return nil, nil
		})
}
