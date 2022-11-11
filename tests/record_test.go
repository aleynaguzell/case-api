package tests

import (
	"case-api/model/record"
	"case-api/storage/repository"
	"testing"
)

func TestGetRecord(t *testing.T) {
	mClient, err := GetMongoClient(t)

	repo := repository.NewRecordsRepository(mClient)

	recordQuery := record.Request{
		"2016-01-26",
		"2018-02-02",
		2700,
		3000,
	}

	records, err := repo.Get(recordQuery)
	if err != nil {
		t.Fail()
	} else if len(records) == 0 {
		t.Fail()
	}
}
