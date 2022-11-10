package repository

import (
	"case-api/model/record"
	"case-api/pkg/config"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var ctx = context.TODO()

type RecordsRepository struct {
	client *mongo.Client
}

type Records interface {
	Get(req record.Request) ([]record.Record, error)
}

func NewRecordsRepository(client *mongo.Client) *RecordsRepository {
	return &RecordsRepository{
		client: client,
	}
}

const collection = "records"

func (r *RecordsRepository) Get(req record.Request) ([]record.Record, error) {

	var results []record.Record

	startDate, err := time.Parse("2006-01-02", req.StartDate)
	if err != nil {
		return nil, err
	}

	endDate, err := time.Parse("2006-01-02", req.EndDate)
	if err != nil {

		return nil, err
	}

	filter := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gt": startDate,
					"$lt": endDate,
				},
			},
		},
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$counts"},
			},
		},
		{
			"$match": bson.M{
				"totalCount": bson.M{
					"$gt": req.MinCount,
					"$lt": req.MaxCount,
				},
			},
		},
	}

	cursor, err := r.client.Database(config.GetConfig().Mongo.Database).Collection(collection).Aggregate(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var record record.Record
		err := cursor.Decode(&record)
		if err != nil {
			return nil, err
		}
		results = append(results, record)
	}

	return results, nil
}
