package tests

import (
	"case-api/pkg/config"
	"case-api/pkg/logger"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

// TODO: change with test mongo url
const testMongoUrl = "mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getircase-study?retryWrites=true"

func init() {
	logger.Init()
	config.Setup("../", "dev")
}

func GetMongoClient(t *testing.T) (*mongo.Client, error) {
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(testMongoUrl))
	if err != nil {
		t.Fail()
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		t.Fail()
	}

	return client, nil
}
