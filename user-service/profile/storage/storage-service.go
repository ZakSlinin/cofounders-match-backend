package storage

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/google/uuid"
	"mime/multipart"
)

type StorageService struct {
	client *s3.Client
	bucket string
}

func NewStorageService(s3Client *s3.Client, bucket string) *StorageService {
	return &StorageService{client: s3Client, bucket: bucket}
}

func (service *StorageService) Upload(ctx context.Context, file multipart.File, header *multipart.FileHeader) (string, error) {
	key := "avatars/" + uuid.New().String() + ".jpg"

	_, err := service.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket:      aws.String(service.bucket),
		Key:         aws.String(key),
		Body:        file,
		ContentType: aws.String(header.Header.Get("Content-Type")),
	})
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("https://storage.yandexcloud.net/%s/%s", service.bucket, key), nil
}
