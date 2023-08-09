package model

import (
	"context"
)

// Article holds article details
type Article struct {
	ID      string `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

type ArticleStore interface {
	Save(ctx context.Context, data *Article) error
	Get(ctx context.Context, id string) (*Article, error)
}
