package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// DynamoDBClient defines the interface for interacting with AWS DynamoDB operations.
type DynamoDBClient interface {
	PutItem(ctx context.Context, tableName string, item map[string]types.AttributeValue) error
	GetItem(ctx context.Context, tableName string, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error)
	DeleteItem(ctx context.Context, tableName string, key map[string]types.AttributeValue) error
	Query(ctx context.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error)
	Scan(ctx context.Context, input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error)
}

type dynamoDBClient struct {
	client *dynamodb.Client
}

// NewDynamoDBClient creates a new DynamoDBClient instance using the provided AWS configuration.
// It initializes a DynamoDB client with the given configuration and returns an interface
// that provides common DynamoDB operations like put, get, delete, query, and scan.
//
// Parameters:
//   - awsCfg: AWS configuration object containing credentials and region settings
//
// Returns:
//   - DynamoDBClient: An interface implementation for DynamoDB operations
func NewDynamoDBClient(awsCfg aws.Config) DynamoDBClient {
	return &dynamoDBClient{
		client: dynamodb.NewFromConfig(awsCfg),
	}
}

// PutItem adds or updates an item in a DynamoDB table.
//
// Parameters:
//   - ctx: Context for the operation
//   - tableName: The name of the DynamoDB table
//   - item: Map of attribute name to AttributeValue for all attributes of the item
//
// Returns:
//   - error: nil if successful, error otherwise
func (d *dynamoDBClient) PutItem(ctx context.Context, tableName string, item map[string]types.AttributeValue) error {
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      item,
	}

	_, err := d.client.PutItem(ctx, input)
	return err
}

// GetItem retrieves an item from a DynamoDB table.
//
// Parameters:
//   - ctx: Context for the operation
//   - tableName: The name of the DynamoDB table
//   - key: Map of attribute name to AttributeValue for the key attributes
//
// Returns:
//   - map[string]types.AttributeValue: The retrieved item's attributes
//   - error: nil if successful, error otherwise
func (d *dynamoDBClient) GetItem(ctx context.Context, tableName string, key map[string]types.AttributeValue) (map[string]types.AttributeValue, error) {
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	}

	result, err := d.client.GetItem(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.Item, nil
}

// DeleteItem removes an item from a DynamoDB table.
//
// Parameters:
//   - ctx: Context for the operation
//   - tableName: The name of the DynamoDB table
//   - key: Map of attribute name to AttributeValue for the key attributes
//
// Returns:
//   - error: nil if successful, error otherwise
func (d *dynamoDBClient) DeleteItem(ctx context.Context, tableName string, key map[string]types.AttributeValue) error {
	input := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       key,
	}

	_, err := d.client.DeleteItem(ctx, input)
	return err
}

// Query executes a query operation on a DynamoDB table.
//
// Parameters:
//   - ctx: Context for the operation
//   - input: The complete QueryInput for the operation
//
// Returns:
//   - *dynamodb.QueryOutput: The query results
//   - error: nil if successful, error otherwise
func (d *dynamoDBClient) Query(ctx context.Context, input *dynamodb.QueryInput) (*dynamodb.QueryOutput, error) {
	return d.client.Query(ctx, input)
}

// Scan performs a scan operation on a DynamoDB table.
//
// Parameters:
//   - ctx: Context for the operation
//   - input: The complete ScanInput for the operation
//
// Returns:
//   - *dynamodb.ScanOutput: The scan results
//   - error: nil if successful, error otherwise
func (d *dynamoDBClient) Scan(ctx context.Context, input *dynamodb.ScanInput) (*dynamodb.ScanOutput, error) {
	return d.client.Scan(ctx, input)
}
