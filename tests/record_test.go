package tests

import (
	"case-api/model/record"
	"case-api/storage/repository"
	"context"
	"testing"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func TestGetRecord(t *testing.T) {
	ctx := context.Background()
	mClient, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://challengeUser:WUMglwNBaydH8Yvu@challenge-xzwqd.mongodb.net/getircase-study?retryWrites=true"))
	if err != nil {
		t.Fail()
	}

	err = mClient.Ping(ctx, nil)
	if err != nil {
		t.Fail()
	}

	defer mClient.Disconnect(context.TODO())

	repo := repository.NewRecordsRepository(mClient)

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
