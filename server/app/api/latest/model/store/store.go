package store

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/obonobo/express-vpn-updater/server/config"
)

var logger = config.Get().Logger()

type Store struct {
	packageName *string
	bucket      *string
	c           *s3.S3
}

func NewStore(bucket string) *Store {
	return &Store{
		bucket: aws.String(bucket),
	}
}

func (s *Store) Get() (string, error) {
	if pkgName := s.pkgName(); pkgName != "" {
		return s.createPkgLink(pkgName), nil
	} else {
		return "", errors.New("not found")
	}
}

func (s *Store) Put(downloadFromUrl string) error {
	if s.alreadyExists(downloadFromUrl) {
		return nil
	}
	file, err := s.downloadFile(downloadFromUrl)
	if err != nil {
		return err
	}
	return s.uploadFile(downloadFromUrl, file)
}

func (s *Store) pkgName() string {
	if s.packageName == nil {
		if name, err := s.getPkgName(); err == nil {
			s.packageName = &name
		} else {
			return ""
		}
	}
	return *s.packageName
}

func (s *Store) alreadyExists(url string) bool {
	return strings.EqualFold(path.Base(url), s.pkgName())
}

func (s *Store) client() *s3.S3 {
	if s.c == nil {
		s.c = s3.New(session.Must(session.NewSession()))
	}
	return s.c
}

func (s *Store) uploadFile(url string, file []byte) error {
	_, err := s.client().PutObject(&s3.PutObjectInput{
		Bucket: s.bucket,
		Key:    aws.String(path.Base(url)),
		Body:   bytes.NewReader(file),
	})
	return err
}

func (s *Store) downloadFile(url string) ([]byte, error) {
	got, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if exceedsMaxFileSize(got) {
		return nil, NewFileTooBigError(got)
	}
	defer got.Body.Close()
	return ioutil.ReadAll(got.Body)
}

func (s *Store) createPkgLink(key string) string {
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", *s.bucket, key)
}

func (s *Store) getPkgName() (key string, err error) {
	logger.Println("Inside Store.getPkgName")
	logger.Println("And the bucket is:", *s.bucket)
	resp, err := s.client().ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: s.bucket,
	})
	logger.Println("In between")
	if err != nil {
		logger.Println("S3 error")
		return "", err
	}
	if len(resp.Contents) <= 0 {
		logger.Println("Content Length is too small")
		return "", errors.New("failure retrieving .deb package from S3 object storage")
	}
	logger.Println(resp, resp.Contents, resp.Contents[0])
	item := resp.Contents[0]
	return *item.Key, nil
}
