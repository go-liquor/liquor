package aws

import "go.uber.org/fx"

// AwsClientModule is a module that provides all the AWS clients.
var AwsClientModule = fx.Module("liquor-module-aws", fx.Provide(
	AwsClient,
	NewS3Client,
	NewSQSClient,
	NewKmsClient,
	NewDynamoDBClient,
))
