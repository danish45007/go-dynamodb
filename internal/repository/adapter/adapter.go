package adapter

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/aws/aws-sdk-go/service/dynamodb/expression"
)

type Database struct {
	connection *dynamodb.DynamoDB
	logMode    bool
}

type Interface interface {
	DbHealth() bool
	FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error)
	FindOne(condition map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error)
	CreateOrUpdate(entity map[string]interface{}, tableName string) (*dynamodb.PutItemOutput, error)
	Delete(condition map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error)
}

func NewAdapter(conn *dynamodb.DynamoDB) Interface {
	return &Database{
		connection: conn,
		logMode:    false,
	}
}

func (db *Database) DbHealth() bool {
	_, err := db.connection.ListTables(&dynamodb.ListTablesInput{})
	if err != nil {
		return false
	}
	return true
}

func (db *Database) FindAll(condition expression.Expression, tableName string) (response *dynamodb.ScanOutput, err error) {
	input := &dynamodb.ScanInput{
		ExpressionAttributeNames:  condition.Names(),
		ExpressionAttributeValues: condition.Values(),
		FilterExpression:          condition.Filter(),
		ProjectionExpression:      condition.Projection(),
		TableName:                 aws.String(tableName),
	}
	db.connection.ScanPages(input, func(page *dynamodb.ScanOutput, lastPage bool) bool {
		return true
	})
	return
}

func (db *Database) FindOne(condition map[string]interface{}, tableName string) (*dynamodb.GetItemOutput, error) {
	parsedCondition, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		fmt.Printf("Error while marshalling %f", err)
		return nil, err
	}
	input := &dynamodb.GetItemInput{
		TableName: aws.String(tableName),
		Key:       parsedCondition,
	}
	response, err := db.connection.GetItem(input)
	if err != nil {
		fmt.Printf("Error while getting item from db %f", err)
		return nil, err
	}
	return response, nil
}

func (db *Database) CreateOrUpdate(entity map[string]interface{}, tableName string) (*dynamodb.PutItemOutput, error) {
	parsedEntity, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		fmt.Printf("Error while marshalling %f", err)
		return nil, err
	}
	input := &dynamodb.PutItemInput{
		TableName: aws.String(tableName),
		Item:      parsedEntity,
	}
	response, err := db.connection.PutItem(input)
	return response, nil
}

func (db *Database) Delete(condition map[string]interface{}, tableName string) (*dynamodb.DeleteItemOutput, error) {
	parsedCondition, err := dynamodbattribute.MarshalMap(condition)
	if err != nil {
		fmt.Printf("Error while marshalling %f", err)
		return nil, err
	}
	item := &dynamodb.DeleteItemInput{
		TableName: aws.String(tableName),
		Key:       parsedCondition,
	}
	response, err := db.connection.DeleteItem(item)
	if err != nil {
		fmt.Printf("Error while deleting item from db %f", err)
		return nil, err
	}
	return response, nil
}
