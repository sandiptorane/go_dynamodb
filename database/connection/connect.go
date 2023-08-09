package connection

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

// GetConnection : returns db connection
func GetConnection() (*dynamodb.Client, error) {
	awsCfg, err := NewAWS()
	if err != nil {
		log.Println("new Aws error", err)
		return nil, err
	}

	return dynamodb.NewFromConfig(awsCfg), nil
}

func NewAWS() (aws.Config, error) {
	awsCfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return aws.Config{}, err
	}

	return awsCfg, nil
}
