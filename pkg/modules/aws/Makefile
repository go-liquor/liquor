.PHONY: genmocks
genmocks:
	mockgen -source=s3.go -package=aws -destination=s3_mock.go
	mockgen -source=kms.go -package=aws -destination=kms_mock.go
	mockgen -source=dynamodb.go -package=aws -destination=dynamodb_mock.go
	mockgen -source=sqs.go -package=aws -destination=sqs_mock.go

