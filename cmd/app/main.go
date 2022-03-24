package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Luis-Miguel-BL/go-dynamodb-crud/config"
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/http/routes"
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/infra/repositories/storedynamodb"
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/internal/migrations"
	"github.com/Luis-Miguel-BL/go-dynamodb-crud/utils/logger"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {
	configs := config.GetConfig()

	connection := storedynamodb.GetConnection()

	emailScoreRepository := storedynamodb.NewEmailScoreDynamoRepo(connection)
	healthRepository := storedynamodb.NewHealthDynamoRepo(connection)

	logger.INFO("Waiting service starting.... ", nil)

	// params := &dynamodb.DeleteTableInput{
	// 	TableName: aws.String("email_score"),
	// }
	// _, err := connection.DeleteTable(params)
	// fmt.Println(err)

	errors := Migrate(connection)
	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("Error on migrate: ", err)
		}
	}
	logger.PANIC("", checkTables(connection))

	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouters(emailScoreRepository, healthRepository)
	logger.INFO("Service running on port ", port)

	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	var errors []error

	callMigrateAndAppendError(&errors, connection, &migrations.Migration{})

	return errors
}

func callMigrateAndAppendError(errors *[]error, connection *dynamodb.DynamoDB, migration migrations.Migrations) {
	err := migration.Migrate(connection)
	if err != nil {
		*errors = append(*errors, err)
	}
}

func checkTables(connection *dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if response != nil {
		if len(response.TableNames) == 0 {
			logger.INFO("Tables not found: ", nil)
		}
		for _, tableName := range response.TableNames {
			logger.INFO("Table found: ", *tableName)
		}
	}
	return err
}
