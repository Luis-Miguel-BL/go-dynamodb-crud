package storedynamodb

import (
	"errors"

	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/entities"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"

	"github.com/fatih/structs"
)

type EmailScoreDynamoRepo struct {
	db        *dynamodb.DynamoDB
	tableName string
}

func NewEmailScoreDynamoRepo(db *dynamodb.DynamoDB) *EmailScoreDynamoRepo {
	return &EmailScoreDynamoRepo{db: db, tableName: "email_score"}
}

func parseDynamoItemToStruct(response map[string]*dynamodb.AttributeValue) (emailScore entities.EmailScore, err error) {
	if response == nil || (response != nil && len(response) == 0) {
		return emailScore, errors.New("Item not found")
	}
	err = dynamodbattribute.UnmarshalMap(response, &emailScore)

	return emailScore, err
}

func parseDynamoItemsToStruct(response []map[string]*dynamodb.AttributeValue) (emailsScore []entities.EmailScore, err error) {

	err = dynamodbattribute.UnmarshalListOfMaps(response, &emailsScore)

	return emailsScore, err
}

func (repository *EmailScoreDynamoRepo) GetByEmail(email string) (emailScore entities.EmailScore, err error) {
	conditionParsed, err := dynamodbattribute.MarshalMap(map[string]interface{}{"email": email})
	if err != nil {
		return emailScore, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(repository.tableName),
		Key:       conditionParsed,
	}
	dynamoOutput, err := repository.db.GetItem(input)
	if err != nil {
		return emailScore, err
	}

	return parseDynamoItemToStruct(dynamoOutput.Item)

}

func (repository *EmailScoreDynamoRepo) ListByEmails(emails []string) (emailsScore []entities.EmailScore, err error) {

	emailsScore = []entities.EmailScore{}
	countEmails := len(emails)

	for i := 0.0; i <= float64(countEmails); i += 99 {
		start := int(i)
		end := int(i + 99)
		if end > countEmails {
			end = countEmails
		}

		results, err := repository.listOneHundredEmails(emails[start:end])
		if err != nil {
			return emailsScore, err
		}
		emailsScore = append(emailsScore, results...)
	}

	return emailsScore, nil

}
func (repository *EmailScoreDynamoRepo) listOneHundredEmails(emails []string) (emailsScore []entities.EmailScore, err error) {
	filterEmails := []expression.OperandBuilder{}

	for _, filterEmail := range emails {
		filterEmails = append(filterEmails, expression.Value(filterEmail))
	}

	filter := expression.In(expression.Name("email"), expression.Value(""), filterEmails...)

	condition, err := expression.NewBuilder().WithFilter(filter).Build()
	if err != nil {
		return emailsScore, err
	}

	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(repository.tableName),
	}
	dynamoOutputs, err := repository.db.Scan(input)

	if err != nil {
		return emailsScore, err
	}

	emailsScore, err = parseDynamoItemsToStruct(dynamoOutputs.Items)
	return emailsScore, err
}

func (repository *EmailScoreDynamoRepo) CreateOrUpdate(entity interface{}, tableName string) (response *dynamodb.PutItemOutput, err error) {
	entityParsed, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		Item:      entityParsed,
		TableName: aws.String(tableName),
	}
	return repository.db.PutItem(input)
}

func (repository *EmailScoreDynamoRepo) Create(entity *entities.EmailScore) (id string, err error) {
	_, err = repository.CreateOrUpdate(structs.Map(entity), repository.tableName)
	return entity.ID, err
}

func (repository *EmailScoreDynamoRepo) Update(entity *entities.EmailScore) error {
	_, err := repository.CreateOrUpdate(structs.Map(entity), repository.tableName)
	return err
}
