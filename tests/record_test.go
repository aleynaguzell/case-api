package tests

import (
	"case-api/model/record"
	"case-api/storage/repository"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestGetRecord(t *testing.T) {
	mClient, err := GetMongoClient(t)
	defer mClient.Disconnect(context.TODO())
	if err != nil {
		t.Fail()
	}
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(testMongoUrl))

	repo := repository.NewRecordsRepository(client)

	recordQuery := record.Request{
		StartDate: "2016-01-26",
		EndDate:   "2018-02-02",
		MinCount:  1900,
		MaxCount:  2700,
	}

	records, err := repo.Get(recordQuery)
	if err != nil {
		t.Fail()
	} else if len(records) == 0 {
		t.Fail()
	}
}