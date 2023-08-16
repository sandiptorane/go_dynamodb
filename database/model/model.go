package model

import (
	"context"
)

// Article holds article details
type Article struct {
	Title   string `json:"title" dynamodbav:"title"`
	Content string `json:"content" dynamodbav:"content"`
	Author  string `json:"author" dynamodbav:"author"`
}

type ArticleStore interface {
	CreateTable(ctx context.Context) error
	DescribeTable(ctx context.Context) error
	Save(ctx context.Context, data *Article) error
	Get(ctx context.Context, title, author string) (*Article, error)
	GetAll(ctx context.Context) ([]*Article, error)
	Update(ctx context.Context, data *Article) error
	Delete(ctx context.Context, data *Article) error
}
