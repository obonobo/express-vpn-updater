package mocks

import (
	"github.com/aws/aws-sdk-go/service/s3"
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

func (msc *MockS3Client) WithListObjectsV2(listObjectsV2 ListObjectsV2) *MockS3Client {
	msc.listObjectsV2 = listObjectsV2
	return msc
}

func (msc *MockS3Client) WithPutObject(putObject PutObject) *MockS3Client {
	msc.putObject = putObject
	return msc
}

func (msc *MockS3Client) ListObjectsV2(input *s3.ListObjectsV2Input) (*s3.ListObjectsV2Output, error) {
	msc.ListObjectInputs = append(msc.ListObjectInputs, input)
	return msc.listObjectsV2(input)
}

func (msc *MockS3Client) PutObject(input *s3.PutObjectInput) (*s3.PutObjectOutput, error) {
	msc.PutObjectInputs = append(msc.PutObjectInputs, input)
	return msc.putObject(input)
}
