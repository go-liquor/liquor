package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/aws/aws-sdk-go-v2/service/sqs/types"
)

// SQSClient defines the interface for interacting with AWS SQS operations.
type SQSClient interface {
	SendMessage(ctx context.Context, queueURL string, messageBody string) (*string, error)
	ReceiveMessages(ctx context.Context, queueURL string, maxMessages int32) ([]types.Message, error)
	DeleteMessage(ctx context.Context, queueURL string, receiptHandle string) error
	GetQueueURL(ctx context.Context, queueName string) (*string, error)
	SendBatchMessages(ctx context.Context, queueURL string, messages []string) error
}

type sqsClient struct {
	client *sqs.Client
}

// NewSQSClient creates a new SQSClient instance using the provided AWS configuration.
// It initializes an SQS client with the given configuration and returns an interface
// that provides common SQS operations like send, receive, delete messages and queue URL retrieval.
//
// Parameters:
//   - awsCfg: AWS configuration object containing credentials and region settings
//
// Returns:
//   - SQSClient: An interface implementation for SQS operations
func NewSQSClient(awsCfg aws.Config) SQSClient {
	return &sqsClient{
		client: sqs.NewFromConfig(awsCfg),
	}
}

// SendMessage sends a single message to an SQS queue.
//
// Parameters:
//   - ctx: Context for the operation
//   - queueURL: The URL of the SQS queue
//   - messageBody: The message content to send
//
// Returns:
//   - *string: The MessageId of the sent message
//   - error: nil if successful, error otherwise
func (s *sqsClient) SendMessage(ctx context.Context, queueURL string, messageBody string) (*string, error) {
	input := &sqs.SendMessageInput{
		QueueUrl:    aws.String(queueURL),
		MessageBody: aws.String(messageBody),
	}

	result, err := s.client.SendMessage(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.MessageId, nil
}

// ReceiveMessages retrieves messages from an SQS queue.
//
// Parameters:
//   - ctx: Context for the operation
//   - queueURL: The URL of the SQS queue
//   - maxMessages: Maximum number of messages to receive (1-10)
//
// Returns:
//   - []types.Message: Array of received messages
//   - error: nil if successful, error otherwise
func (s *sqsClient) ReceiveMessages(ctx context.Context, queueURL string, maxMessages int32) ([]types.Message, error) {
	input := &sqs.ReceiveMessageInput{
		QueueUrl:            aws.String(queueURL),
		MaxNumberOfMessages: maxMessages,
	}

	result, err := s.client.ReceiveMessage(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.Messages, nil
}

// DeleteMessage removes a message from the queue after processing.
//
// Parameters:
//   - ctx: Context for the operation
//   - queueURL: The URL of the SQS queue
//   - receiptHandle: The receipt handle of the message to delete
//
// Returns:
//   - error: nil if successful, error otherwise
func (s *sqsClient) DeleteMessage(ctx context.Context, queueURL string, receiptHandle string) error {
	input := &sqs.DeleteMessageInput{
		QueueUrl:      aws.String(queueURL),
		ReceiptHandle: aws.String(receiptHandle),
	}

	_, err := s.client.DeleteMessage(ctx, input)
	return err
}

// GetQueueURL retrieves the URL of an SQS queue by its name.
//
// Parameters:
//   - ctx: Context for the operation
//   - queueName: The name of the SQS queue
//
// Returns:
//   - *string: The URL of the queue
//   - error: nil if successful, error otherwise
func (s *sqsClient) GetQueueURL(ctx context.Context, queueName string) (*string, error) {
	input := &sqs.GetQueueUrlInput{
		QueueName: aws.String(queueName),
	}

	result, err := s.client.GetQueueUrl(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.QueueUrl, nil
}

// SendBatchMessages sends multiple messages to an SQS queue in a single request.
//
// Parameters:
//   - ctx: Context for the operation
//   - queueURL: The URL of the SQS queue
//   - messages: Array of message bodies to send
//
// Returns:
//   - error: nil if successful, error otherwise
func (s *sqsClient) SendBatchMessages(ctx context.Context, queueURL string, messages []string) error {
	var entries []types.SendMessageBatchRequestEntry
	for i, msg := range messages {
		entries = append(entries, types.SendMessageBatchRequestEntry{
			Id:          aws.String(string(rune(i))),
			MessageBody: aws.String(msg),
		})
	}

	input := &sqs.SendMessageBatchInput{
		QueueUrl: aws.String(queueURL),
		Entries:  entries,
	}

	_, err := s.client.SendMessageBatch(ctx, input)
	return err
}
