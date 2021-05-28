package mocks

import (
	"testing"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/obonobo/express-vpn-updater/server/app/util/testutils"
	"github.com/stretchr/testify/assert"
)

const (
	once = testutils.Once
)

var (
	defaulMockListObjectsV2Output = &s3.ListObjectsV2Output{}
	defaultMockPutObjectOutput    = &s3.PutObjectOutput{}
)

// A mocked version of the S3Client that the Store will use as backing storage.
// The mock records the inputs that it was given.
type MockS3Client struct {
	ListObjectInputs []*s3.ListObjectsV2Input
	PutObjectInputs  []*s3.PutObjectInput
	listObjectsV2    ListObjectsV2
	putObject        PutObject
}
type ListObjectsV2 func(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error)
type PutObject func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error)

func NewMockS3Client() *MockS3Client {
	return &MockS3Client{
		[]*s3.ListObjectsV2Input{},
		[]*s3.PutObjectInput{},
		func(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
			return defaulMockListObjectsV2Output, nil
		},
		func(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
			return defaultMockPutObjectOutput, nil
		},
	}
}

func (msc *MockS3Client) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	msc.ListObjectInputs = append(msc.ListObjectInputs, input)
	return msc.listObjectsV2(input)
}

func (msc *MockS3Client) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	msc.PutObjectInputs = append(msc.PutObjectInputs, input)
	return msc.putObject(input)
}

func (msc *MockS3Client) WithListObjectsV2(listObjectsV2 ListObjectsV2) *MockS3Client {
	msc.listObjectsV2 = listObjectsV2
	return msc
}

func (msc *MockS3Client) WithPutObject(putObject PutObject) *MockS3Client {
	msc.putObject = putObject
	return msc
}

func (msc *MockS3Client) AssertBucketListed(
	t *testing.T,
	bucketName string,
	msgAndArgs ...interface{},
) (this *MockS3Client) {
	msc.AssertBucketListedMultipleTimes(t, bucketName, once, msgAndArgs...)
	return msc
}

func (msc *MockS3Client) AssertBucketListedMultipleTimes(
	t *testing.T,
	bucketName string,
	numberOfTimes int,
	msgAndArgs ...interface{},
) (this *MockS3Client) {
	numberOfTimesItWasCalled := 0
	for _, v := range msc.ListObjectInputs {
		if *v.Bucket == bucketName {
			numberOfTimesItWasCalled++
		}
	}
	assert.Equal(t, numberOfTimes, numberOfTimesItWasCalled, msgAndArgs...)
	return msc
}

func (msc *MockS3Client) AssertPutObjectWasCalled(
	t *testing.T,
	bucketName string,
	msgAndArgs ...interface{},
) (this *MockS3Client) {
	msc.AssertPutObjectWasCalledMultipleTimes(t, bucketName, once, msgAndArgs...)
	return msc
}

func (msc *MockS3Client) AssertPutObjectWasCalledMultipleTimes(
	t *testing.T,
	bucketName string,
	numberOfTimes int,
	msgAndArgs ...interface{},
) (this *MockS3Client) {
	numberOfTimesItWasCalled := 0
	for _, v := range msc.PutObjectInputs {
		if *v.Bucket == bucketName {
			numberOfTimesItWasCalled++
		}
	}
	assert.Equal(t, numberOfTimes, numberOfTimesItWasCalled, msgAndArgs...)
	return msc
}

// Runs a custom assertion on the most recent recorded PutObject input. If no
// inputs exist, then the assertion fails.
func (msc *MockS3Client) AssertPutObjectInput(
	t *testing.T,
	assertion func(input *s3.PutObjectInput),
) (this *MockS3Client) {
	assert.NotEmpty(t, msc.PutObjectInputs, "there must be at least 1 recorded PutObject input")
	assertion(msc.PutObjectInputs[len(msc.PutObjectInputs)-1])
	return msc
}

// Runs a custom assertion on the most recent recored ListObjectsV2 input. If no
// input exists, then the assertion fails.
func (msc *MockS3Client) AssertListObjectsV2Input(
	t *testing.T,
	assertion func(input *s3.ListObjectsV2Input),
) (this *MockS3Client) {
	assert.NotEmpty(t, msc.ListObjectInputs, "there must be at least 1 recoreded ListObjectsV2 input")
	assertion(msc.ListObjectInputs[len(msc.ListObjectInputs)-1])
	return msc
}
