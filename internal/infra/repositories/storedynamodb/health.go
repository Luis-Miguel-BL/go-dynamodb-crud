package storedynamodb

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type HealthDynamoRepo struct {
	db *dynamodb.DynamoDB
}

func NewHealthDynamoRepo(db *dynamodb.DynamoDB) *HealthDynamoRepo {
	return &HealthDynamoRepo{db: db}
}

func (repository *HealthDynamoRepo) Health() bool {
	_, err := repository.db.ListTables(&dynamodb.ListTablesInput{})

	return err == nil
}
