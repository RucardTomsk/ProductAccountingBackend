package s3

import (
	"context"
	"github.com/minio/minio-go/v7"
	"io"
	"net/http"
	"net/url"
	"productAccounting-v1/internal/domain/base"
	"strings"
	"time"
)

type UploadInput struct {
	File        io.Reader
	Name        string
	Size        int64
	ContentType string
}

// MinioService communicates with minio (s3).
type MinioService struct {
	client   *minio.Client
	bucket   string
	endpoint string
}

func NewMinioService(client *minio.Client, bucket, endpoint string) *MinioService {
	return &MinioService{
		client:   client,
		bucket:   bucket,
		endpoint: endpoint,
	}
}

func (s *MinioService) Upload(ctx context.Context, input UploadInput) *base.ServiceError {
	opts := minio.PutObjectOptions{
		ContentType:  input.ContentType,
		UserMetadata: map[string]string{"x-amz-acl": "public-read"},
	}

	name := input.Name + "." + strings.Split(input.ContentType, "/")[1]
	_, err := s.client.PutObject(ctx, s.bucket, name, input.File, input.Size, opts)
	if err != nil {
		return unexpectedServiceError(err)
	}

	return nil
}

func (s *MinioService) GetFileURL(ctx context.Context, fileName string) (*string, *base.ServiceError) {
	reqParams := make(url.Values)
	reqParams.Set("response-content-disposition", "attachment; filename=\""+fileName+"\"")

	_, err := s.client.StatObject(context.Background(), s.bucket, fileName+".png", minio.StatObjectOptions{})
	if err != nil {
		_, err := s.client.StatObject(context.Background(), s.bucket, fileName+".jpeg", minio.StatObjectOptions{})
		if err != nil {
			return nil, base.NewNotFoundError(err)
		} else {
			fileName += ".jpeg"
		}
	} else {
		fileName += ".png"
	}
	resignedURL, err := s.client.PresignedGetObject(ctx, s.bucket, fileName, time.Duration(1000)*time.Second, reqParams)
	if err != nil {
		return nil, unexpectedServiceError(err)
	}

	urlString := resignedURL.String()
	return &urlString, nil
}

// unexpectedServiceError returns any unclassified service error.
func unexpectedServiceError(err error) *base.ServiceError {
	return &base.ServiceError{
		Err:     err,
		Blame:   base.BlameServer,
		Code:    http.StatusInternalServerError,
		Message: "unexpected error occurred",
	}
}
