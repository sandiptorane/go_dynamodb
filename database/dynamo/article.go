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

func (a *ArticleStore) DescribeTable(ctx context.Context) error {
	param := &dynamodb.DescribeTableInput{TableName: aws.String("article")}
	_, err := a.db.DescribeTable(ctx, param)
	if err != nil {
		log.Println("describe table error", err)
		return err
	}

	return nil
}

// CreateTable create table
func (a *ArticleStore) CreateTable(ctx context.Context) error {
	input := dynamodb.CreateTableInput{
		TableName: aws.String("article"),
		AttributeDefinitions: []types.AttributeDefinition{
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
				AttributeName: aws.String("title"),
				KeyType:       types.KeyTypeHash,
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

	log.Println("input item", items)

	input := &dynamodb.PutItemInput{
		Item:      items,
		TableName: aws.String("article"),
	}

	_, err = a.db.PutItem(ctx, input)

	return err
}

// Get Article
func (a *ArticleStore) Get(ctx context.Context, title, author string) (*model.Article, error) {
	value, err := attributevalue.MarshalMap(map[string]string{
		"title":  title,
		"author": author,
	})
	if err != nil {
		return nil, err
	}

	input := &dynamodb.GetItemInput{
		Key:       value,
		TableName: aws.String("article"),
	}

	res, err := a.db.GetItem(ctx, input)
	if err != nil {
		log.Println("error fetching article", err, "input", input.Key)
		return nil, err
	}

	if res.Item == nil {
		return nil, nil
	}

	var article model.Article

	err = attributevalue.UnmarshalMap(res.Item, &article)
	if err != nil {
		log.Println("unmarshall error", err)
		return nil, err
	}

	return &article, nil
}

func (a *ArticleStore) GetAll(ctx context.Context) ([]*model.Article, error) {
	input := &dynamodb.ScanInput{
		TableName: aws.String("article"),
	}

	res, err := a.db.Scan(ctx, input)
	if err != nil {
		log.Println("error fetching article", err)
		return nil, err
	}

	var articles []*model.Article

	for _, a := range res.Items {
		var article model.Article
		err = attributevalue.UnmarshalMap(a, &article)
		if err != nil {
			log.Println("unmarshall error", err)
			return nil, err
		}

		articles = append(articles, &article)
	}

	return articles, nil
}

func (a *ArticleStore) Update(ctx context.Context, data *model.Article) error {
	input := &dynamodb.UpdateItemInput{
		ExpressionAttributeValues: map[string]types.AttributeValue{
			":content": &types.AttributeValueMemberS{Value: data.Content},
		},

		UpdateExpression: aws.String("set content = :content"),
		Key: map[string]types.AttributeValue{
			"title":  &types.AttributeValueMemberS{Value: data.Title},
			"author": &types.AttributeValueMemberS{Value: data.Author},
		},

		TableName: aws.String("article"),
	}

	_, err := a.db.UpdateItem(ctx, input)

	return err
}

func (a *ArticleStore) Delete(ctx context.Context, data *model.Article) error {
	input := &dynamodb.DeleteItemInput{
		Key: map[string]types.AttributeValue{
			"title":  &types.AttributeValueMemberS{Value: data.Title},
			"author": &types.AttributeValueMemberS{Value: data.Author},
		},

		TableName: aws.String("article"),
	}

	_, err := a.db.DeleteItem(ctx, input)

	return err
}
