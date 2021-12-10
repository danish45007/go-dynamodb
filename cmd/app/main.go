package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/danish45007/go-dynamodb/config"
	"github.com/danish45007/go-dynamodb/internal/repository/adapter"
	"github.com/danish45007/go-dynamodb/internal/repository/instance"
	"github.com/danish45007/go-dynamodb/internal/routes"
	"github.com/danish45007/go-dynamodb/utils/logger"
)

func main() {
	configs := config.GetConfig()
	connection := instance.GetConnection()
	repository := adapter.NewAdapter(connection)
	logger.INFO("waiting for the service to start...", nil)
	errors := Migrate(connection)
	if len(errors) > 0 {
		for _, err := range errors {
			logger.PANIC("error while database migration", err)
		}
	}
	logger.PANIC("", checkTables(connection))
	port := fmt.Sprintf(":%v", configs.Port)
	router := routes.NewRouter().SetRouter(repository)
	logger.INFO("service is running on port", port)
	server := http.ListenAndServe(port, router)
	log.Fatal(server)
}

func Migrate(connection *dynamodb.DynamoDB) []error {
	var errors []error
	callMigrateAndAppendError
}

func callMigrateAndAppendError(errors) {}

func checkTables(connection *dynamodb.DynamoDB) error {
	response, err := connection.ListTables(&dynamodb.ListTablesInput{})
	if response != nil {
		if len(response.TableNames) == 0 {
			logger.INFO("tables not found :(", nil)
		}
		for _, table := range response.TableNames {
			logger.INFO("table found:", &table)
		}
	}
	return err
}
