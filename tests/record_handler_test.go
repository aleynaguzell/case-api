package tests

import (
	"bytes"
	"case-api/api/handler"
	"case-api/services"
	"case-api/storage/repository"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestRecordController(t *testing.T) {

	expected := []byte(`{"code":0,"msg":"Success","records":[{"key":"TAKwGc6Jr4i8Z487","createdAt":"2017-01-28T01:22:14.398Z","totalCount":0},{"key":"NAeQ8eX7e5TEg7oH","createdAt":"2017-01-27T08:19:14.135Z","totalCount":0}]}`)

	postBody := []byte(`{
		"startDate": "2016-01-26",
		"endDate": "2018-02-02",
		"minCount": 2700,
		"maxCount": 3000
	}`)

	req := httptest.NewRequest(http.MethodPost, "/records", bytes.NewBuffer(postBody))
	mClient, err := GetMongoClient(t)

	recordRepository := repository.NewRecordsRepository(mClient)
	recordService := services.NewRecordService(recordRepository)
	recordHandler := handler.NewRecordHandler(*recordService)

	w := httptest.NewRecorder()

	handler := http.HandlerFunc(recordHandler.Get)

	handler.ServeHTTP(w, req)
	res := w.Result()
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		t.Fatal(err)
	}

	if http.StatusOK != res.StatusCode {
		t.Error("expected", http.StatusOK, "got status", res.StatusCode)
	}
	if !strings.Contains(string(body), string(expected)) {
		t.Error("expected"+string(expected)+"got", string(body))
	}
}
