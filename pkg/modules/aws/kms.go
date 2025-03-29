package aws

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/kms"
	"github.com/aws/aws-sdk-go-v2/service/kms/types"
)

// KMSClient defines the interface for interacting with AWS KMS operations.
type KMSClient interface {
	Encrypt(ctx context.Context, keyID string, plaintext []byte) ([]byte, error)
	Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error)
	GenerateDataKey(ctx context.Context, keyID string, keySpec string) (*DataKey, error)
	GenerateRandom(ctx context.Context, numberOfBytes int32) ([]byte, error)
}

// DataKey represents a data key with both plaintext and encrypted versions
type DataKey struct {
	Plaintext      []byte
	EncryptedBytes []byte
}

type kmsClient struct {
	client *kms.Client
}

// NewKmsClient creates a new KMSClient instance using the provided AWS configuration.
// It initializes a KMS client with the given configuration and returns an interface
// that provides common KMS operations like encrypt, decrypt, generate data keys,
// and generate random bytes.
//
// Parameters:
//   - awsCfg: AWS configuration object containing credentials and region settings
//
// Returns:
//   - KMSClient: An interface implementation for KMS operations
func NewKmsClient(awsCfg aws.Config) KMSClient {
	return &kmsClient{
		client: kms.NewFromConfig(awsCfg),
	}
}

// Encrypt encrypts plaintext using the specified KMS key.
//
// Parameters:
//   - ctx: Context for the operation
//   - keyID: The ID or ARN of the KMS key
//   - plaintext: The data to encrypt
//
// Returns:
//   - []byte: The encrypted data
//   - error: nil if successful, error otherwise
func (k *kmsClient) Encrypt(ctx context.Context, keyID string, plaintext []byte) ([]byte, error) {
	input := &kms.EncryptInput{
		KeyId:     aws.String(keyID),
		Plaintext: plaintext,
	}

	result, err := k.client.Encrypt(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.CiphertextBlob, nil
}

// Decrypt decrypts ciphertext that was encrypted using a KMS key.
//
// Parameters:
//   - ctx: Context for the operation
//   - ciphertext: The data to decrypt
//
// Returns:
//   - []byte: The decrypted plaintext
//   - error: nil if successful, error otherwise
func (k *kmsClient) Decrypt(ctx context.Context, ciphertext []byte) ([]byte, error) {
	input := &kms.DecryptInput{
		CiphertextBlob: ciphertext,
	}

	result, err := k.client.Decrypt(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.Plaintext, nil
}

// GenerateDataKey generates a unique data key for client-side encryption.
//
// Parameters:
//   - ctx: Context for the operation
//   - keyID: The ID or ARN of the KMS key
//   - keySpec: The length and type of data key (e.g., "AES_256")
//
// Returns:
//   - *DataKey: Contains both plaintext and encrypted versions of the data key
//   - error: nil if successful, error otherwise
func (k *kmsClient) GenerateDataKey(ctx context.Context, keyID string, keySpec string) (*DataKey, error) {
	input := &kms.GenerateDataKeyInput{
		KeyId:   aws.String(keyID),
		KeySpec: types.DataKeySpec(keySpec),
	}

	result, err := k.client.GenerateDataKey(ctx, input)
	if err != nil {
		return nil, err
	}

	return &DataKey{
		Plaintext:      result.Plaintext,
		EncryptedBytes: result.CiphertextBlob,
	}, nil
}

// GenerateRandom generates a secure random byte string using KMS.
//
// Parameters:
//   - ctx: Context for the operation
//   - numberOfBytes: The number of bytes of random data to generate
//
// Returns:
//   - []byte: The generated random data
//   - error: nil if successful, error otherwise
func (k *kmsClient) GenerateRandom(ctx context.Context, numberOfBytes int32) ([]byte, error) {
	input := &kms.GenerateRandomInput{
		NumberOfBytes: aws.Int32(numberOfBytes),
	}

	result, err := k.client.GenerateRandom(ctx, input)
	if err != nil {
		return nil, err
	}

	return result.Plaintext, nil
}
