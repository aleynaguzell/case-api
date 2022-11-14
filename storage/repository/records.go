package repository

import (
	"case-api/model/record"
	"case-api/pkg/config"
	"context"

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

	filter := []bson.M{
		{
			"$match": bson.M{
				"createdAt": bson.M{
					"$gt": req.StartDate,
					"$lt": req.EndDate,
				},
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
		{
			"$project": bson.M{
				"_id":        0,
				"key":        1,
				"createdAt":  1,
				"totalCount": bson.M{"$sum": "$totalCount"},
			},
		},
	}

	cursor, err := r.client.Database(config.GetConfig().Mongo.Database).Collection(collection).Aggregate(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var record record.Record
		err := cursor.Decode(&record)
		if err != nil {
			return nil, err
		}
		results = append(results, record)
	}

	return results, nil
}
