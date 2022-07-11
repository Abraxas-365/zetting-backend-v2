package internal_S3

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func (s *S3Client) Upload(filename multipart.File, keyName string) (*s3manager.UploadOutput, error) {
	uploader := s3manager.NewUploader(s.Sess)
	// Upload the file to S3.
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(keyName),
		Body:   filename,
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}
