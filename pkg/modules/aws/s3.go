package aws

import (
	"context"
	"io"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

// NewS3Client creates a new S3Client instance using the provided AWS configuration.
// It initializes an S3 client with the given configuration and returns an interface
// that provides common S3 operations like upload, download, delete, list, and URL signing.
//
// Parameters:
//   - awsCfg: AWS configuration object containing credentials and region settings
//
// Returns:
//   - S3Client: An interface implementation for S3 operations
func NewS3Client(awsCfg aws.Config) S3Client {
	return &s3client{
		client: s3.NewFromConfig(awsCfg),
	}
}

// S3Client defines the interface for interacting with AWS S3 operations.
type S3Client interface {
	UploadFile(ctx context.Context, bucket, key string, file io.Reader) error
	DownloadFile(ctx context.Context, bucket, key string) ([]byte, error)
	DeleteFile(ctx context.Context, bucket, key string) error
	ListFiles(ctx context.Context, bucket, prefix string) ([]string, error)
	GetSignedURL(ctx context.Context, bucket, key string, expires time.Duration) (string, error)
}

type s3client struct {
	client *s3.Client
}

// UploadFile uploads a file to an S3 bucket.
//
// Parameters:
//   - ctx: Context for the operation
//   - bucket: The name of the S3 bucket
//   - key: The object key (path) in the bucket
//   - file: The file content as an io.Reader
//
// Returns:
//   - error: nil if successful, error otherwise
func (s *s3client) UploadFile(ctx context.Context, bucket, key string, file io.Reader) error {
	_, err := s.client.PutObject(ctx, &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   file,
	})
	return err
}

// DownloadFile retrieves a file from an S3 bucket.
//
// Parameters:
//   - ctx: Context for the operation
//   - bucket: The name of the S3 bucket
//   - key: The object key (path) in the bucket
//
// Returns:
//   - []byte: The file contents
//   - error: nil if successful, error otherwise
func (s *s3client) DownloadFile(ctx context.Context, bucket, key string) ([]byte, error) {
	result, err := s.client.GetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	defer result.Body.Close()

	return io.ReadAll(result.Body)
}

// DeleteFile removes a file from an S3 bucket.
//
// Parameters:
//   - ctx: Context for the operation
//   - bucket: The name of the S3 bucket
//   - key: The object key (path) to delete
//
// Returns:
//   - error: nil if successful, error otherwise
func (s *s3client) DeleteFile(ctx context.Context, bucket, key string) error {
	_, err := s.client.DeleteObject(ctx, &s3.DeleteObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	return err
}

// ListFiles retrieves a list of files from an S3 bucket with an optional prefix.
//
// Parameters:
//   - ctx: Context for the operation
//   - bucket: The name of the S3 bucket
//   - prefix: The prefix to filter objects (optional)
//
// Returns:
//   - []string: List of object keys in the bucket
//   - error: nil if successful, error otherwise
func (s *s3client) ListFiles(ctx context.Context, bucket, prefix string) ([]string, error) {
	var files []string
	paginator := s3.NewListObjectsV2Paginator(s.client, &s3.ListObjectsV2Input{
		Bucket: aws.String(bucket),
		Prefix: aws.String(prefix),
	})

	for paginator.HasMorePages() {
		page, err := paginator.NextPage(ctx)
		if err != nil {
			return nil, err
		}

		for _, obj := range page.Contents {
			files = append(files, *obj.Key)
		}
	}

	return files, nil
}

// GetSignedURL generates a pre-signed URL for temporary access to an S3 object.
//
// Parameters:
//   - ctx: Context for the operation
//   - bucket: The name of the S3 bucket
//   - key: The object key (path) in the bucket
//   - expires: The duration until the URL expires
//
// Returns:
//   - string: The pre-signed URL
//   - error: nil if successful, error otherwise
func (s *s3client) GetSignedURL(ctx context.Context, bucket, key string, expires time.Duration) (string, error) {
	presignClient := s3.NewPresignClient(s.client)

	request, err := presignClient.PresignGetObject(ctx, &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	}, s3.WithPresignExpires(expires))

	if err != nil {
		return "", err
	}

	return request.URL, nil
}
