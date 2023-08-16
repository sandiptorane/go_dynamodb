package dynamo

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/stretchr/testify/assert"
	"go_dynamodb/database/connection"
	"log"
	"testing"
)

func TestArticleStore_CreateTable(t *testing.T) {
	type fields struct {
		db *dynamodb.Client
	}

	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr error
	}{
		{
			name:   "Happy path",
			fields: fields{},
			args: args{
				context.TODO(),
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, err := connection.GetConnection()
			if err != nil {
				log.Println("connection error:", err)
				return
			}

			a := &ArticleStore{
				db: conn,
			}

			err = a.CreateTable(tt.args.ctx)
			if tt.wantErr != nil {
				assert.EqualError(t, err, tt.wantErr.Error())
				return
			}

			assert.Nil(t, err)
		})
	}
}
