package dynamodb

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go/aws/request"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// Person is a simple struct that we will use to store data in DynamoDB.
type Person struct {
	Name string
}

type ddbClient interface {
	// PutItemWithContext fakes method to the interface for putting an item into DynamoDB. Look at how the client is used in the `Save` method below.
	PutItemWithContext(
		ctx context.Context,
		input *dynamodb.PutItemInput,
		opts ...request.Option,
	) (*dynamodb.PutItemOutput, error)
}

// DynamoDBSaver is a wrapper that encapsulates interactions with DynamoDB.
type DynamoDBSaver struct {
	Client ddbClient
}

// Save performs the put operation on DynamoDB using the client we've assigned to DynamoDBSaver.
func (s *DynamoDBSaver) Save(ctx context.Context, p *Person) error {
	item, err := dynamodbattribute.MarshalMap(p)
	if err != nil {
		return fmt.Errorf("failed to marshal shoutout for storage: %s", err)
	}

	input := &dynamodb.PutItemInput{
		Item:      item,
		TableName: aws.String(os.Getenv("TABLE_NAME")),
	}

	_, err = s.Client.PutItemWithContext(ctx, input)

	return err
}
