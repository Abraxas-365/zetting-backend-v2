package ports

import (
	"bytes"
	"mime/multipart"

	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type BucketRepo interface {
	Upload(filename multipart.File, keyName string) (*s3manager.UploadOutput, error)
	Get(key string) (*bytes.Buffer, error)
}
