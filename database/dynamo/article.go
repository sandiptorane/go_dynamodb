package dynamo

import (
	"context"
	"go_dynamodb/database/model"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type ArticleStore struct {
	db *dynamodb.Client
}

// GetInstance returns instance of articleStore
func GetInstance(db *dynamodb.Client) *ArticleStore {
	return &ArticleStore{db: db}
}

func (a *ArticleStore) CreateTable(ctx context.Context) error {
	input := dynamodb.CreateTableInput{
		TableName: aws.String("article"),
		AttributeDefinitions: []types.AttributeDefinition{
			{
				AttributeName: aws.String("id"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("title"),
				AttributeType: types.ScalarAttributeTypeS,
			},
			{
				AttributeName: aws.String("author"),
				AttributeType: types.ScalarAttributeTypeS,
			}},
		KeySchema: []types.KeySchemaElement{
			{
				AttributeName: aws.String("id"),
				KeyType:       types.KeyTypeHash,
			},
			{
				AttributeName: aws.String("title"),
				KeyType:       types.KeyTypeRange,
			},
			{
				AttributeName: aws.String("author"),
				KeyType:       types.KeyTypeRange,
			},
		},
		ProvisionedThroughput: &types.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(10),
			WriteCapacityUnits: aws.Int64(5),
		},
	}

	_, err := a.db.CreateTable(ctx, &input)
	if err != nil {
		log.Println("error creating table", err)
		return err
	}

	return nil
}

// Save article
func (a *ArticleStore) Save(ctx context.Context, data *model.Article) error {
	items, err := attributevalue.MarshalMap(data)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      items,
		TableName: aws.String("article"),
	}

	_, err = a.db.PutItem(ctx, input)

	return err
}

// Get Article
func (a *ArticleStore) Get(ctx context.Context, id string) (*model.Article, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},

		TableName: aws.String("article"),
	}

	res, err := a.db.GetItem(ctx, input)
	if err != nil {
		log.Println("error fetching article", err)
		return nil, err
	}

	var article model.Article

	err = attributevalue.UnmarshalMap(res.Item, &article)
	if err != nil {
		log.Println("unmarshall error", err)
		return nil, err
	}

	return &article, nil
}
