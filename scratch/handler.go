package main

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/obonobo/express-vpn-updater/util"
)

func main() {
	lambda.Start(Scratch)
}

func Scratch(req util.Request) (util.Response, error) {

	stringified, err := json.Marshal(listS3())
	if err != nil {
		return util.Response{}, err
	}

	return util.BasicMessage(string(stringified)), nil
}

func listS3() []string {
	sess := session.Must(session.NewSession())
	client := s3.New(sess)

	// Use the client to list the items in a bucket
	items, err := client.ListObjectsV2(&s3.ListObjectsV2Input{
		Bucket:  aws.String("express-vpn-deb-cache"),
		MaxKeys: aws.Int64(100),
	})

	if err != nil {
		panic(err)
	}

	contents := items.Contents
	mapped := []string{}

	for _, v := range contents {
		mapped = append(mapped, *v.Key)
	}

	return mapped
}
