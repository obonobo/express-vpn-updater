package store

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"path"
	"testing"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/obonobo/express-vpn-updater/server/app/util/testutils"
	"github.com/obonobo/express-vpn-updater/server/app/util/testutils/mocks"
	"github.com/stretchr/testify/assert"
)

const (
	testBucket   = "test-bucket"
	testFileBody = "My hot body"
	once         = testutils.Once
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

	expectedGetResponse = fmt.Sprintf(
		PackageLinkFormat,
		testBucket,
		*mockListObjectsV2Output.Contents[0].Key)
)

// Tests S3Store.Get()
func TestGet(t *testing.T) {
	var (
		mockHttpClient = createMockHttpClient()
		mockS3Client   = createMockS3Client()
		store          = New(mockS3Client, mockHttpClient, testBucket)
	)

	resp, err := store.Get()
	if err != nil {
		assert.Failf(t, "store.Get should not return an error", "%v", err)
	}
	assert.Equal(t, expectedGetResponse, resp)
}

// Tests S3Store.Put()
func TestPut(t *testing.T) {
	var (
		mockDownloadUrl    = "http://tests.com/is/this/thing/on/???"
		mockHttpClient     = createMockHttpClient()
		mockS3Client       = createMockS3Client()
		store              = New(mockS3Client, mockHttpClient, testBucket)
		failOnErrorMessage = "placing an item into the store should not fail"
	)

	if err := store.Put(mockDownloadUrl); err != nil {
		assert.Failf(t, failOnErrorMessage, "%v", err)
	}

	mockHttpClient.AssertUrlPinged(t, mockDownloadUrl, once)
	mockS3Client.
		AssertBucketListed(t, testBucket).
		AssertListObjectsV2Input(t, func(input *s3.ListObjectsV2Input) {
			assert.Equal(t, testBucket, *input.Bucket)
		}).
		AssertPutObjectWasCalled(t, testBucket).
		AssertPutObjectInput(t, func(input *s3.PutObjectInput) {
			bod, err := ioutil.ReadAll(input.Body)
			if err != nil {
				assert.Failf(t, "PutObject input body should be readable, but was unable to read it", "%v", err)
			}
			assert.Equal(t, testFileBody, string(bod))
			assert.Equal(t, testBucket, *input.Bucket)
			assert.Equal(t, path.Base(mockDownloadUrl), *input.Key)
		})
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
			return &http.Response{
				Status:        "200 OK",
				Body:          io.NopCloser(bytes.NewBufferString(testFileBody)),
				ContentLength: 666,
			}, nil
		})
}
