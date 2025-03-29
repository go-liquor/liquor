package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/go-liquor/liquor/v2/config"
	"go.uber.org/zap"
)

// AwsClient creates and returns an AWS configuration using the provided Config and logger.
// It loads the default AWS configuration with the region specified in the config.
// If the configuration loading fails, it logs a fatal error using the provided logger.
//
// Parameters:
//   - config: A pointer to the Config struct containing AWS configuration settings
//   - logger: A zap logger instance for error logging
//
// Returns:
//   - aws.Config: The AWS configuration object
func AwsClient(config *config.Config, logger *zap.Logger) aws.Config {
	cfg, err := awsconfig.LoadDefaultConfig(context.TODO(),
		awsconfig.WithRegion(config.GetString("aws.region")),
		awsconfig.WithCredentialsProvider(aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     config.GetString("aws.accessKeyId"),
				SecretAccessKey: config.GetString("aws.secretAccessKey"),
			}, nil
		})))
	if err != nil {
		logger.Fatal("failed to load aws config", zap.Error(err))
	}
	return cfg
}
