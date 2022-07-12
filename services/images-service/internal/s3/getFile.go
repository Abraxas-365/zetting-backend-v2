package internal_S3

import (
	"bytes"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *S3Client) Get(keyName string) (*bytes.Buffer, error) {
	result, err := s.Svc.GetObject(
		&s3.GetObjectInput{
			Bucket: aws.String(s.Bucket),
			Key:    aws.String(keyName),
		})
	if err != nil {
		return nil, err
	}
	buf := new(bytes.Buffer)
	buf.ReadFrom(result.Body)
	return buf, nil
}
