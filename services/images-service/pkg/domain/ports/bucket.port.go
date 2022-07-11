package ports

import (
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type BucketRepo interface {
	Upload(filename multipart.File, keyName string) (*s3manager.UploadOutput, error)
}
