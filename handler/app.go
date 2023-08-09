package handler

import (
	"go_dynamodb/database/dynamo"
	"go_dynamodb/database/model"

	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Application struct {
	ArticleStore model.ArticleStore
}

func NewApplication(db *dynamodb.Client) *Application {
	return &Application{
		ArticleStore: dynamo.GetInstance(db),
	}
}
