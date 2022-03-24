package migrations

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Migrations interface {
	Migrate(connection *dynamodb.DynamoDB) error
}
