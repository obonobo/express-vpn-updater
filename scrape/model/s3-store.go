package model

import (
	"io/ioutil"
	"net/http"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	bucket *string = aws.String("express-vpn-deb-cache")
	key    *string = aws.String("my-package.deb")
)

type Store struct {
	client *s3.S3
}

func NewPackageStore() *Store {
	return &Store{}
}

func (s *Store) UploadPackageFrom(url string) (*string, error) {

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	up, err := s.getClient().PutObject(&s3.PutObjectInput{
		Bucket: bucket,
		Key:    key,
	})

	if err != nil {
		return nil, err
	}

	ret := up.String()
	return &ret, nil
}

func (s *Store) DownloadPackage() (*string, error) {
	pkg, err := s.getClient().GetObject(&s3.GetObjectInput{
		Bucket: bucket,
		Key:    key,
	})
	if err != nil {
		return nil, err
	}
	defer pkg.Body.Close()
	downloadMe, err := ioutil.ReadAll(pkg.Body)
	if err != nil {
		return nil, err
	}

	ret := string(downloadMe)
	return &ret, nil
}

func (s *Store) getClient() *s3.S3 {
	if s.client == nil {
		s.initClient()
	}
	return s.client
}

func (s *Store) initClient() {
	s.client = s3.New(session.Must(session.NewSession()))
}
