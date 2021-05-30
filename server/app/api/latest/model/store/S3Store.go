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
	"github.com/obonobo/express-vpn-updater/server/app/config"
	"github.com/obonobo/express-vpn-updater/server/app/util"
	"github.com/obonobo/express-vpn-updater/server/app/util/logging"
)

var (
	logger            *logging.Logger = config.Get().Logger()
	defaultHttpClient util.HttpClient = http.DefaultClient
	defaultS3Client   S3Client        = nil
)

type S3Store struct {
	sc          S3Client
	hc          util.HttpClient
	bucket      *string
	packageName *string
}

func New(s3client S3Client, httpClient util.HttpClient, bucket string) Store {
	return &S3Store{
		bucket: aws.String(bucket),
		sc:     s3client,
		hc:     httpClient,
	}
}

func NewStoreWithBucket(bucket string) Store {
	return New(defaultS3Client, defaultHttpClient, bucket)
}

func Default() Store {
	return NewStoreWithBucket(config.Get().Bucket)
}

func (s *S3Store) Get() (string, error) {
	if pkgName := s.pkgName(); pkgName != "" {
		return s.createPkgLink(pkgName), nil
	} else {
		return "", errors.New("not found")
	}
}

func (s *S3Store) Put(downloadFromUrl string) error {
	if s.alreadyExists(downloadFromUrl) {
		return nil
	}
	file, err := s.downloadFile(downloadFromUrl)
	if err != nil {
		return err
	}
	return s.uploadFile(downloadFromUrl, file)
}

func (s *S3Store) pkgName() string {
	if s.packageName == nil {
		if name, err := s.getPkgName(); err == nil {
			s.packageName = &name
		} else {
			return ""
		}
	}
	return *s.packageName
}

func (s *S3Store) alreadyExists(url string) bool {
	return strings.EqualFold(path.Base(url), s.pkgName())
}

func (s *S3Store) client() S3Client {
	if s.sc == nil {
		s.sc = s3.New(session.Must(session.NewSession()))
	}
	return s.sc
}

func (s *S3Store) uploadFile(url string, file []byte) error {
	_, err := s.client().PutObject(&s3.PutObjectInput{
		Bucket: s.bucket,
		Key:    aws.String(path.Base(url)),
		Body:   bytes.NewReader(file),
	})
	return err
}

func (s *S3Store) downloadFile(url string) ([]byte, error) {
	got, err := s.hc.Get(url)
	if err != nil {
		return nil, err
	}
	if exceedsMaxFileSize(got) {
		return nil, NewFileTooBigError(got)
	}
	defer got.Body.Close()
	return ioutil.ReadAll(got.Body)
}

func (s *S3Store) createPkgLink(key string) string {
	return fmt.Sprintf("https://%s.s3.amazonaws.com/%s", *s.bucket, key)
}

func (s *S3Store) getPkgName() (key string, err error) {
	logger.Inside("S3Store.getPkgName")
	logger.Println("And the bucket is:", *s.bucket)
	resp, err := s.client().ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket: s.bucket,
	})
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
