# aws


```go
package main

import (
    "github.com/go-liquor/liquor/pkg/modules/aws" // add this
)

func main() {
	app.NewApp(
        aws.AwsClientModule, // add this
	)
}
```

## AWS Services

### DynamoDB Client

```go
type DynamoDBClient interface {
    // PutItem adds or updates an item in a DynamoDB table
    PutItem(ctx context.Context, tableName string, item map[string]types.AttributeValue) error
    
    // GetItem retrieves an item from a DynamoDB table using its key
    GetItem(ctx context.Context, tableName string, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error)
    
    // DeleteItem removes an item from a DynamoDB table
    DeleteItem(ctx context.Context, tableName string, key map[string]types.AttributeValue) error
    
    // Query executes a query operation on a DynamoDB table
    Query(ctx context.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
    
    // Scan performs a scan operation on a DynamoDB table
    Scan(ctx context.Context, input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
}
```

### KMS Client

```go
type KMSClient interface {
    // Encrypt encrypts plaintext using the specified KMS key
    Encrypt(ctx context.Context, keyID string, plaintext []byte) ([]byte, error)
    
    // Decrypt decrypts ciphertext that was encrypted using a KMS key
    Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error)
    
    // GenerateDataKey generates a unique data key for client-side encryption
    GenerateDataKey(ctx context.Context, keyID string, keySpec string) (*DataKey, error)
    
    // GenerateRandom generates a secure random byte string using KMS
    GenerateRandom(ctx context.Context, numberOfBytes int32) ([]byte, error)
}
```

### S3 Client

```go
type S3Client interface {
    // UploadFile uploads a file to an S3 bucket
    UploadFile(ctx context.Context, bucket, key string, file io.Reader) error
    
    // DownloadFile downloads a file from an S3 bucket
    DownloadFile(ctx context.Context, bucket, key string) ([]byte, error)
    
    // DeleteFile removes a file from an S3 bucket
    DeleteFile(ctx context.Context, bucket, key string) error
    
    // ListFiles lists files in a bucket with an optional prefix
    ListFiles(ctx context.Context, bucket, prefix string) ([]string, error)
    
    // GetSignedURL generates a pre-signed URL for temporary access to a file
    GetSignedURL(ctx context.Context, bucket, key string, expires time.Duration) (string, error)
}
```

### SQS Client

```go
type SQSClient interface {
    // SendMessage sends a single message to an SQS queue
    SendMessage(ctx context.Context, queueURL string, messageBody string) (*string, error)
    
    // ReceiveMessages retrieves messages from an SQS queue
    ReceiveMessages(ctx context.Context, queueURL string, maxMessages int32) ([]types.Message, error)
    
    // DeleteMessage removes a message from the queue after processing
    DeleteMessage(ctx context.Context, queueURL string, receiptHandle string) error
    
    // GetQueueURL retrieves the URL of an SQS queue by its name
    GetQueueURL(ctx context.Context, queueName string) (*string, error)
    
    // SendBatchMessages sends multiple messages to an SQS queue in a single request
    SendBatchMessages(ctx context.Context, queueURL string, messages []string) error
}
```

## Usage Example

In your service, you can inject and use any of these clients:

```go
type Service struct {
    dynamoDB aws.DynamoDBClient
    kms     aws.KMSClient
    s3      aws.S3Client
    sqs     aws.SQSClient
}

func NewService(
    dynamoDB aws.DynamoDBClient,
    kms aws.KMSClient,
    s3 aws.S3Client,
    sqs aws.SQSClient,
) *Service {
    return &Service{
        dynamoDB: dynamoDB,
        kms:     kms,
        s3:      s3,
        sqs:     sqs,
    }
}
```

All AWS clients will be automatically injected by the framework.

## Mock

For testing purposes, you can use gomock to create mock implementations. Here's how to use them:

```go
// In your test file
func TestYourService(t *testing.T) {
    // Create gomock controller
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    // Create mock instances
    mockS3 := mocks.NewMockS3Client(ctrl)
    mockDynamoDB := mocks.NewMockDynamoDBClient(ctrl)
    mockKMS := mocks.NewMockKMSClient(ctrl)
    mockSQS := mocks.NewMockSQSClient(ctrl)

    // Configure mock behaviors
    mockS3.EXPECT().
        UploadFile(gomock.Any(), "bucket-name", "key", gomock.Any()).
        Return(nil)
    
    mockDynamoDB.EXPECT().
        GetItem(gomock.Any(), "table-name", gomock.Any()).
        Return(map[string]types.AttributeValue{
            "id": &types.AttributeValueMemberS{Value: "123"},
        }, nil)

    // Initialize your service with mocks
    service := NewService(
        mockDynamoDB,
        mockKMS,
        mockS3,
        mockSQS,
    )

    // Run your tests
    err := service.YourMethod()
    require.NoError(t, err)
}

