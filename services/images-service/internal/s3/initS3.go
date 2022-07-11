package internal_S3

import (
	"errors"
	"service-images/pkg/domain/ports"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	Bucket string
	Region string
	Sess   *session.Session
	Svc    *s3.S3
}

func NewS3Client(id string, seccret string, bucket string, region string) (ports.BucketRepo, error) {
	// Initialize a session in us-west-2 that the SDK will use to load
	creds := credentials.NewStaticCredentials(id, seccret, "")
	sess, err := session.NewSession(&aws.Config{
		Credentials: creds,
		Region:      aws.String(region)},
	)
	if err != nil {
		return nil, errors.New("repository.s3Bucket")
	}
	// Create S3 service client
	return &S3Client{
		Bucket: bucket,
		Region: region,
		Sess:   sess,
		Svc:    s3.New(sess),
	}, nil

}
